package main

import (
	"fmt"
	"strconv"
)

func main() {
	graphs := []string{"graphs/custom_format/graph_1.xml",
		"graphs/custom_format/graph_2.xml",
		"graphs/custom_format/graph_3.xml",
		"graphs/custom_format/graph_4.xml"}

	csv_file := "log.csv"
	traces := get_traces(csv_file, ";")

	results := make([][][]string, len(graphs))

	for i, filename := range graphs {
		graph := create_DCR_graph(filename)
		results[i] = make([][]string, 2)
		results[i] = graph.Evaluate_traces(traces)
	}
	var output string
	for i, result := range results {
		output += fmt.Sprintln("Graph_" + strconv.Itoa(i+1) + ":")
		output += fmt.Sprint("Correct:" + strconv.Itoa(len(result[0])) + " - ")
		output += fmt.Sprintln("Incorrect:" + strconv.Itoa(len(result[1])))
	}
	fmt.Println(output)
}
