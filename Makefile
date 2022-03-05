SHELL =/bin/bash
dir = $(shell pwd)

build:
	go build -o ./bin/main ./src
all: build
	./bin/main
run:
	go run ./src
echo_host:
	echo "http://localhost:6060"
doc: echo_host
	godoc -goroot ${dir}
current_dir :
	${SHELL} $$(pwd)