package main

import "strings"

func create_DCR_graph(file string) DCR_graph {

	graph := new(DCR_graph)
	dcr := decode_xml(file)

	graph.filename = file
	graph.nodes = dcr.Events.Names
	graph.marking = dcr.Markings
	graph.conditions_for = extract_constraints(dcr.Constraints.Conditions)
	graph.responses_to = extract_constraints(dcr.Constraints.Responses)
	graph.excludes_to = extract_constraints(dcr.Constraints.Excludes)
	graph.includes_to = extract_constraints(dcr.Constraints.Includes)

	return *graph
}

func extract_constraints(constraints []string) Constraint_map {
	c_map := new(Constraint_map)
	c_map.constraint_map = make(map[string][]string)

	for i := 0; i < len(constraints); i++ {
		var splits = strings.Split(constraints[i], ":")

		if _, ok := c_map.constraint_map[splits[0]]; ok {
			c_map.constraint_map[splits[0]] = append(c_map.constraint_map[splits[0]], splits[1])
		} else {
			c_map.constraint_map[splits[0]] = []string{splits[1]}
		}
	}

	return *c_map
}

func get_all_relations_to(relations_for Constraint_map, event string) []string {
	keys := make([]string, 0)
	for k, value := range relations_for.constraint_map {
		if string_slice_contains(value, event) {
			keys = append(keys, k)
		}
	}
	return keys
}

func remove_all_excluded(exclude Constraint_map, included []string, event string) []string {
	result := included
	for _, to_exclude := range exclude.constraint_map[event] {
		index_to_delete := find_index(result, to_exclude)
		if index_to_delete != -1 {
			result = remove_index(result, index_to_delete)
		}
	}
	return result
}
