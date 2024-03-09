build:
	@go mod download
	@go build -o web_server .
run:
	@go mod download
	@go run .
test:
	@go mod download
	go test ./... 
