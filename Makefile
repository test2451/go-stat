build:
ifeq ($(OS),Windows_NT)
	go build -o build/pie-statas.exe ./main.go
else
	go build -o build/pie-statas ./main.go
endif

install:
ifeq ($(OS),Windows_NT)
	go install ./main.go
else
	go install ./main.go
endif

.PHONY: build install
