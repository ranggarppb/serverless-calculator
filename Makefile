ifndef $(GOPATH)
    GOPATH=$(shell go env GOPATH)
    export GOPATH
endif

local_engine:
	export FUNCTION_TARGET=Calculate && go run $(shell pwd)/cmd/main.go

mockery-gen:
	@rm -rf ./mocks
	$(GOPATH)/bin/mockery --name ICalculatorService --dir ./types
