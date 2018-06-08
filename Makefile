.PHONY: all bars
GOPATH0 := $(firstword $(subst :, ,$(GOPATH)))

all: bars

bars:
	go build -o $(GOPATH0)/bin/bars.so -buildmode=plugin ./bars
