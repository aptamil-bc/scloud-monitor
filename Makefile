build:
ifeq ($(OS),Windows_NT)
	go build -o build/scloud-monitor.exe main.go
else
	go build -o build/scloud-monitor main.go
endif

install:
ifeq ($(OS),Windows_NT)
	go install main.go
else
	go install main.go
endif

.PHONY: build install
