package main

import (
    "fmt"
    "encoding/xml"
    "os"
    "strings"
)

type DCR struct {
    Events Event `xml:"events"`
    Constraints Constraint `xml:"constraints"`
    Markings Marking `xml:"markings"`
}

type Event struct {
    Names []string `xml:"event"`
}

type Constraint struct {
    Conditions []string `xml:"conditions>condition"`
    Responses []string `xml:"responses>response"`
    Corresponses []string `xml:"corresponses>corresponse"`
    Excludes []string `xml:"excludes>exclude"`
    Includes []string `xml:"includes>include"`
    Milestones []string `xml:"milestones>milestone"`
    Spawns []string `xml:"spawns>spawn"`
}

type Marking struct {
    Executed []string `xml:"executed>event"`
    Included []string `xml:"included>event"`
    Pending []string `xml:"pending>event"`
}

type Constraint_map struct {
     constraint_map map[string][]string
}

type DCR_graph struct{
    nodes []string

    conditions_for Constraint_map
    milestones_for Constraint_map
    responses_to   Constraint_map
    excludes_to    Constraint_map
    includes_to    Constraint_map

    marking Marking
}

func create_DCR_graph(file string) DCR_graph{

    dcr_graph := new(DCR_graph)
    dcr := decode_xml(file)

    dcr_graph.nodes = dcr.Events.Names
    dcr_graph.marking = dcr.Markings
    dcr_graph.conditions_for = extract_constraints(dcr.Constraints.Conditions)

    return *dcr_graph
}

func extract_constraints(constraints []string) Constraint_map{
    c_map := new(Constraint_map)
    c_map.constraint_map = make(map[string][]string)

    for i:=0; i < len(constraints); i++ {
        var splits = strings.Split(constraints[i], ":")

        if _, ok := c_map.constraint_map[splits[0]]; ok {
            c_map.constraint_map[splits[0]] = append(c_map.constraint_map[splits[0]], splits[1])
        } else {
            c_map.constraint_map[splits[0]] = []string{splits[1]}
        }
    }

    return *c_map
}

func decode_xml(file string) DCR {
    xmlFile, err := os.Open(file)

    if err != nil {
        fmt.Println(err)
    }

    fmt.Println("Successfully Opened %s", file)

    // defer the closing of our xmlFile so that we can parse it later on
    defer xmlFile.Close()

    dcr := new(DCR)

    decoded := xml.NewDecoder(xmlFile)
    err = decoded.Decode(dcr)

    if err != nil {
       fmt.Println(err)
    }
    return *dcr
}

func string_slice_contains(input []string, target string) bool {
   for _, a := range input {
      if a == target {
         return true
      }
   }
   return false
}

func get_all_conditions_to(conditions_for Constraint_map, event string) []string{
    keys := make([]string, 0)
    for k, value:= range conditions_for.constraint_map{
        if(string_slice_contains(value, event)){
            keys = append(keys, k)
        }
    }
    return keys
}

func retain_all(included []string, keys []string) []string{
    incon := make([]string, 0)
    for _,con := range keys{
        if(string_slice_contains(included, con)){
            incon = append(incon,con)
        }
    }
    return incon
}

func contains_all(executed []string, incon []string) bool{
    for _, con := range incon{
        if(!string_slice_contains(executed, con)){
            return false
        }
    }
    return true
}

func find_index(strings []string,  event string) int {
    i := 0
    for _, s := range strings{
        if(s == event){
            return i
        }
        i++
    }
    return -1
}

func remove_index(s []string, index int) []string {
	return append(s[:index], s[index+1:]...)
}

func enabled(dcr_graph DCR_graph, event string) bool {

    if(!string_slice_contains(dcr_graph.nodes, event)){return true}

    if(!string_slice_contains(dcr_graph.marking.Included, event)){return false}

    keys := get_all_conditions_to(dcr_graph.conditions_for, event)
    incon := retain_all(dcr_graph.marking.Included, keys)
    contains_all := contains_all(dcr_graph.marking.Executed, incon)
    if(!contains_all){return false}

    return true
}

func execute(dcr_graph DCR_graph, event string) Marking {
    if(!string_slice_contains(dcr_graph.nodes, event)){return dcr_graph.marking}
    if(!enabled(dcr_graph, event)){return  dcr_graph.marking}
    result := dcr_graph.marking
    result.Executed = append(result.Executed, event)
    result.Pending = remove_index(result.Pending, find_index(result.Pending,event))

    fmt.Println(dcr_graph.responses_to.constraint_map[event])
    return result
}

func enabled_test (dcr_graph DCR_graph, events []string){
    for _, event := range events{
        fmt.Println(event, enabled(dcr_graph,event))
    }
}

func main(){
    var filename string = "graph.xml"
    graph := create_DCR_graph(filename)
    //enabled_test(graph, graph.nodes)

    result := execute(graph, "Fill_out_application")
    fmt.Println("\n", result)
}
