package main

type DCR interface {
	Enabled(event string) bool
	Execute(event string) Marking
	Get_included_pending() []string
	Is_accepting() bool
	Enabled_events() []string
	Execute_trace(trace []string) bool
	Evaluate_traces(traces map[string][]string) [][]string
}

func (graph DCR_graph) Enabled(event string) bool {

	if !string_slice_contains(graph.nodes, event) {
		return true
	}
	if !string_slice_contains(graph.marking.Included, event) {
		return false
	}

	keys := get_all_relations_to(graph.conditions_for, event)
	incon := retain_all(graph.marking.Included, keys)
	contains_all := contains_all(graph.marking.Executed, incon)
	if !contains_all {
		return false
	}

	return true
}

func (graph DCR_graph) Execute(event string) Marking {

	if !string_slice_contains(graph.nodes, event) {
		return graph.marking
	}
	if !graph.Enabled(event) {
		return graph.marking
	}

	if !string_slice_contains(graph.marking.Executed, event) {
		graph.marking.Executed = append(graph.marking.Executed, event)
	}

	index_to_delete := find_index(graph.marking.Pending, event)
	if index_to_delete != -1 {
		graph.marking.Pending = remove_index(graph.marking.Pending, index_to_delete)
	}

    for _, event := range graph.responses_to.constraint_map[event]{
        if(!string_slice_contains(graph.marking.Pending, event)){
            graph.marking.Pending = append(graph.marking.Pending, event)
        }
    }

	graph.marking.Included =
		remove_all_excluded(graph.excludes_to, graph.marking.Included, event)

	graph.marking.Included =
		append(graph.marking.Included, graph.includes_to.constraint_map[event]...)

	return graph.marking
}

func (graph DCR_graph) Get_included_pending() []string {

	keys := graph.marking.Included
	result := retain_all(graph.marking.Pending, keys)
	return result
}

func (graph DCR_graph) Is_accepting() bool {

	if len(graph.Get_included_pending()) == 0 {
		return true
	}
	return false
}

func (graph DCR_graph) Enabled_events() []string {
	legal_events := make([]string, 0)

	for _, event := range graph.nodes {
		if graph.Enabled(event) {
			legal_events = append(legal_events, event)
		}
	}
	return legal_events
}

func (graph DCR_graph) Execute_trace(trace []string) bool {
	legal_events := graph.Enabled_events()

	for _, event := range trace {
		if string_slice_contains(legal_events, event) ||
			!string_slice_contains(graph.nodes, event) {
			graph.marking = graph.Execute(event)
		} else {
			return false
		}
		legal_events = graph.Enabled_events()
	}
	return graph.Is_accepting()
}

func (graph DCR_graph) Evaluate_traces(traces map[string][]string) [][]string {
	satisfied_traces := make([]string, 0)
	unsatisfied_traces := make([]string, 0)

	for id, trace := range traces {
		if graph.Execute_trace(trace) {
			satisfied_traces = append(satisfied_traces, id)
		} else {
			unsatisfied_traces = append(unsatisfied_traces, id)
		}

		graph = create_DCR_graph(graph.filename)
	}
	result := make([][]string, 2)
	result[0] = satisfied_traces
	result[1] = unsatisfied_traces
	return result
}
