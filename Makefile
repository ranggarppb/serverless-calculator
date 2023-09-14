ifndef $(GOPATH)
    GOPATH=$(shell go env GOPATH)
    export GOPATH
endif

local_function:
	export FUNCTION_TARGET=Calculate && go run $(shell pwd)/app/main.go function

console:
	go run $(shell pwd)/app/main.go console

mockery-gen:
	@rm -rf ./mocks
	$(GOPATH)/bin/mockery --name ICalculatorService --dir ./types/interfaces
