package main

import "fmt"
import "strconv"

type marking struct {
    executed, included, pending []string
}

type DCRgraph struct{
    nodes []string
    markings marking
}

func create_DCRgraph(number_of_nodes int) DCRgraph {
    g:= DCRgraph{nodes:make([]string, number_of_nodes)}
    g.markings = marking{executed:make([]string, number_of_nodes),
        included:make([]string, number_of_nodes),
        pending:make([]string, number_of_nodes)}
    return g
}

func populated_graph(graph DCRgraph, number_of_nodes int){
    for i := 0; i < number_of_nodes; i++{
        graph.nodes[i] = strconv.Itoa(i)
    }
}

func printGraph(graph DCRgraph){
    fmt.Println(graph.nodes)
}

func main(){
    number_of_nodes := 10
    graph1 := create_DCRgraph(number_of_nodes)
    populated_graph(graph1, number_of_nodes)
    printGraph(graph1)
}

