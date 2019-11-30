.PHONY: standard build run all

CC=go

standard: run

build:
	go build DCR_graph.go DCR_structs.go get_log.go

run:
	go build DCR_graph.go DCR_structs.go get_log.go
	./DCR_graph
test:
	go build test.go get_log.go
	./test
