# go commands
GOCMD=go
GOBUILD=$(GOCMD) build
GOTEST=$(GOCMD) test
GOCLEAN=$(GOCMD) clean
BINARY_NAME=calculator.exe

#Below "make" commands are available
#make build: only build binary
#make clean: clean binary
#make run:   build and run program
#make test:  only to test program
#make all:   build and run

build: 
	$(GOBUILD) -o $(BINARY_NAME) -v
clean:
	$(GOCLEAN)
run: build
	$(BINARY_NAME)
test:
	$(GOTEST) -v ./... -cover
all : run

.PHONY: build test run clean all