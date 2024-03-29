BINARY_NAME = bf
TEST_COMMAND = gotest

.PHONY: build
build:
	go build -v -o $(BINARY_NAME) ./cmd/$(BINARY_NAME)

.PHONY: run
run: build
	./$(BINARY_NAME) 

.PHONY: test
test: 
	$(TEST_COMMAND) -v -cover -parallel 5 -failfast  ./... 

.PHONY: tidy
tidy:
	go mod tidy

# auto restart bot (using fiber CLI <3)
.PHONY: dev
dev:
	fiber dev -t ./cmd/$(BINARY_NAE)
