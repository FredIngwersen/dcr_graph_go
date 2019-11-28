.PHONY: standard build run all

CC=go

standard: run 

build:
	go build DCR_graph.go DCR_structs.go

run: 
	go build DCR_graph.go DCR_structs.go
	./DCR_graph
