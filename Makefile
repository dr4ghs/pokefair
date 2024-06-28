all: test run

test:
	go test -v ./...

build:
	CGO_ENABLED=1 go build -o bin/main ./cmd/pokefair/main.go
	CGO_ENABLED=1 CC=x86_64-w64-mingw32-gcc GOOS=windows GOARCH=amd64 go build -o /mnt/d/Workspace/Games/main.exe -ldflags "-H=windowsgui -s -w" ./cmd/pokefair/main.go

run: build
	./bin/main
	
