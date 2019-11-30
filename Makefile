.PHONY: standard build run test all

CC=go

standard: run

build:
	go build main.go DCR_interface.go DCR_structs.go helper_functions.go get_log.go

run:
	go build main.go DCR_functions.go DCR_interface.go DCR_structs.go helper_functions.go get_log.go
	./main
test:
	go build test.go get_log.go
	./test
