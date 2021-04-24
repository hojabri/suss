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
run:
	./$(BINARY_NAME)
docker-up:
	MODE=dev docker-compose up --remove-orphans --build -d