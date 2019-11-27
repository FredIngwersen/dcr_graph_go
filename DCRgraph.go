package main

import "fmt"
import "strconv"

type relations struct {
    conditions_for map[string][]string
    milestones_for map[string][]string
    responses_to   map[string][]string
    excludes_to    map[string][]string
    includes_to    map[string][]string
}

type DCRgraph struct{
    nodes []string
}

func create_DCRgraph(number_of_nodes int) DCRgraph {
    g:= DCRgraph{name:make([]string, number_of_nodes)}
    return g
}

func populated_graph(graph DCRgraph, number_of_nodes int){
    for i := 0; i < number_of_nodes; i++{
        graph.name[i] = strconv.Itoa(i)
    }
}

func printGraph(graph DCRgraph){
    fmt.Println(graph.name)
}

func main(){
    number_of_nodes := 10
    graph1 := createDCRgraph(number_of_nodes)
    populated_graph(graph1, number_of_nodes)
    printGraph(graph1)
}

