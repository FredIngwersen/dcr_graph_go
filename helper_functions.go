package main

import (
	"encoding/xml"
	"fmt"
	"os"
)

func decode_xml(file string) DCR_XML {
	xmlFile, err := os.Open(file)

	if err != nil {
		fmt.Println(err)
	}

	// defer the closing of our xmlFile so that we can parse it later on
	defer xmlFile.Close()

	dcr := new(DCR_XML)

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

func retain_all(included []string, keys []string) []string {
	incon := make([]string, 0)
	for _, con := range keys {
		if string_slice_contains(included, con) {
			incon = append(incon, con)
		}
	}
	return incon
}

func contains_all(executed []string, incon []string) bool {
	for _, con := range incon {
		if !string_slice_contains(executed, con) {
			return false
		}
	}
	return true
}

func find_index(strings []string, event string) int {
	i := 0
	for _, s := range strings {
		if s == event {
			return i
		}
		i++
	}
	return -1
}

func remove_index(s []string, index int) []string {
	return append(s[:index], s[index+1:]...)
}
