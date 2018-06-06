GOPATH0 := $(firstword $(subst :, ,$(GOPATH)))

all: history

history:
	go build -o $(GOPATH0)/bin/history.so -buildmode=plugin .
