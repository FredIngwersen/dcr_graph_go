package main

import (
    "fmt"
    "encoding/xml"
    "os"
    "strings"
)

func create_DCR_graph(file string) DCR_graph{

    graph := new(DCR_graph)
    dcr := decode_xml(file)

    graph.nodes = dcr.Events.Names
    graph.marking = dcr.Markings
    graph.conditions_for = extract_constraints(dcr.Constraints.Conditions)
    graph.responses_to= extract_constraints(dcr.Constraints.Responses)
    graph.excludes_to = extract_constraints(dcr.Constraints.Excludes)
    graph.includes_to= extract_constraints(dcr.Constraints.Includes )

    return *graph
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

func get_all_relations_to(relations_for Constraint_map, event string) []string{
    keys := make([]string, 0)
    for k, value:= range relations_for.constraint_map{
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


func enabled(graph DCR_graph, event string) bool {

    if(!string_slice_contains(graph.nodes, event)){return true}

    if(!string_slice_contains(graph.marking.Included, event)){return false}

    keys := get_all_relations_to(graph.conditions_for, event)
    incon := retain_all(graph.marking.Included, keys)
    contains_all := contains_all(graph.marking.Executed, incon)
    if(!contains_all){return false}

    return true
}

func remove_all_excluded(exclude Constraint_map,included []string ,event string)[]string{
    result := included
    for _, to_exclude := range exclude.constraint_map[event]{
        index_to_delete := find_index(result,to_exclude)
        if(index_to_delete != -1){
            result= remove_index(result, index_to_delete)
        }
    }
    return result
}

func execute(graph DCR_graph, event string) DCR_graph{
    if(!string_slice_contains(graph.nodes, event)) {return graph}
    if(!enabled(graph, event)) {return  graph}

    result := graph.marking

    if(!string_slice_contains(result.Executed, event)){
        result.Executed = append(result.Executed, event)
    }

    index_to_delete := find_index(result.Pending, event)
    if(index_to_delete != -1){
        result.Pending = remove_index(result.Pending, index_to_delete)
    }

    result.Pending =
        append(result.Pending, graph.responses_to.constraint_map[event]...)

    result.Included =
        remove_all_excluded(graph.excludes_to, result.Included, event)

    result.Included =
        append(result.Included, graph.includes_to.constraint_map[event]...)
    graph.marking = result
    return graph
}

func get_included_pending(graph DCR_graph) []string{

    keys := graph.marking.Included
    result := retain_all(graph.marking.Pending, keys)
    return result
}

func is_accepting(graph DCR_graph) bool {
    if(len(get_included_pending(graph)) == 0){
        return true
    }
    return false
}

func enabled_events(graph DCR_graph) []string {
    legal_events := make([]string, 0)

    for _, event := range graph.nodes {
        if (enabled(graph, event)) {
            legal_events = append(legal_events, event)
        }
    }
    return legal_events
}

func execute_trace(graph DCR_graph, trace []string) bool {
    legal_events := enabled_events(graph)

    for _, event := range trace {
        if (string_slice_contains(legal_events, event)) {
            graph = execute(graph, event)
        } else {
            return false
        }
        legal_events = enabled_events(graph)
    }

    return is_accepting(graph)
}

func evaluate_traces(graph DCR_graph, traces map[string]string) ([]string, []string) {
    satisfied_traces:= make([]string, 0)
    unsatisfied_traces:= make([]string, 0)

    for id, trace := range traces  {
        if (execute_trace(graph, trace)) {
            satisfied_traces = append(satisfied_traces, id)
        } else {
            unsatisfied_traces = append(unsatisfied_traces, id)
        }

        graph = create_DCR_graph(filename)
    }

    return satisfied_traces, unsatisfied_traces
}

func main(){
    graphs := []string{ "graphs/custom_format/graph_1",
                        "graphs/custom_format/graph_2",
                        "graphs/custom_format/graph_3",
                        "graphs/custom_format/graph_4"}

    csv_file := "log.csv"
    traces := get_traces(csv_file, ";")
    fmt.Println(traces[0])
`
    results := make([][]string, len(graphs))

    for i, filename := range graphs{
        graph := create_DCR_graph(filename)
        results[i] = make([]string, 2)
        results[i] = evaluate_traces(graph, traces)
    }
        `

}
