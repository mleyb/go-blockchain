BIN := bin
OUTPUT := go-blockchain

.PHONY : build clean fmt lint test

build-linux:
	echo "Compiling for Linux"
	GOOS=linux go build -v -o ${BIN}/${OUTPUT} .	

build-windows:
	echo "Compiling for Windows X64"
	GOOS=windows GOARCH=amd64 go build -v -o ${BIN}/${OUTPUT}.exe .	

build: dep build-linux build-windows

dep:
	go mod download -x

clean:
	go clean
	rm -rf ${BIN}

lint:
	go get -u golang.org/x/lint/golint
	go vet ./...
	golint -set_exit_status ./...

fmt:
	go fmt ./...

test:
	go test -v -cover ./...