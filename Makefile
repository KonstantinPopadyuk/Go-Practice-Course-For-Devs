TASK ?= 01-slices

.PHONY: check test task race fmt

check:
	go vet ./...
	go test -run '^$$' ./...

test:
	go test ./...

task:
	go test ./tasks/$(TASK) -v

race:
	go test ./tasks/$(TASK) -race -v

fmt:
	gofmt -w $$(find tasks -name '*.go' -type f)
