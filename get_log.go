package main

import (
	"encoding/csv"
	"io"
	"log"
	"os"
	"strings"
)

func get_traces(csv_file string, seperator string) map[string][]string {
	file, _ := os.Open(csv_file)
	r := csv.NewReader(file)

	traces := make(map[string][]string)

	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		split := strings.Split(record[0], seperator)
		id := split[0]
		event := split[5]

		if _, ok := traces[id]; ok {
			traces[id] = append(traces[id], event)
		} else {
			traces[id] = []string{event}
		}
	}

	return traces
}
