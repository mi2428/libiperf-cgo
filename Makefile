.PHONY: build
build: main.go
	@mkdir -p bin
	@go build -o ./bin/iperf3 main.go