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

func main(){
    var filename string = "graph.xml"
    graph := create_DCR_graph(filename)
    fmt.Println(graph.conditions_for)
}
