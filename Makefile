ifndef $(GOPATH)
    GOPATH=$(shell go env GOPATH)
    export GOPATH
endif

build:
	GOFLAGS=-mod=mod go build -o bin/serverless-calculator app/main.go

local_function:
	export FUNCTION_TARGET=Calculate && go run $(shell pwd)/app/main.go function

console:
	go run $(shell pwd)/app/main.go console

mockery-gen:
	@rm -rf ./mocks
	$(GOPATH)/bin/mockery --name ICalculatorService --dir ./objects/calculation
	$(GOPATH)/bin/mockery --name ICalculationInput --dir ./objects/calculation
