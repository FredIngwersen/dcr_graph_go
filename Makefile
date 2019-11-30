.PHONY: standard build run all

CC=go

standard: run 

build:
	go build main.go DCR_interface.go DCR_structs.go helper_functions.go get_log.go

run: 
	go build main.go DCR_interface.go DCR_structs.go helper_functions.go get_log.go
	./DCR_graph
