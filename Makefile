dev:
	air

fmt:
	go fmt ./...

tidy:
	go mod tidy

build:
	go build -o bin/go-portfolio main.go

clean:
	if exist bin (del /Q /S bin\*)
