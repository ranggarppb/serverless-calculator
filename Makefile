ifndef $(GOPATH)
    GOPATH=$(shell go env GOPATH)
    export GOPATH
endif

local_rest_engine:
	export FUNCTION_TARGET=Calculate && go run $(shell pwd)/app/main.go rest

local_console_engine:
	go run $(shell pwd)/app/main.go console

mockery-gen:
	@rm -rf ./mocks
	$(GOPATH)/bin/mockery --name ICalculatorService --dir ./types
