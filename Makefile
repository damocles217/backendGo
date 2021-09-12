BINARY = server
FILE = main.go

main: $(FILE)
	go mod download
	go build -o $(BINARY) $(FILE)

test: $(FILE)
	go test