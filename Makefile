# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOTEST=$(GOCMD) test
WEBSERVICE_BINARY_NAME=susswebservice
WEBSERVICE_MAIN=cmd/webservice/main.go


all: build
build:
	$(GOBUILD) -o $(WEBSERVICE_BINARY_NAME) -v $(WEBSERVICE_MAIN)
test:
	$(GOTEST) -v ./...
clean:
	rm -f $(BINARY_NAME)
run:
	./$(BINARY_NAME)