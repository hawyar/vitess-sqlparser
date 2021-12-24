build:
	GOOS=linux GOARCH=amd64 go build -o bin/linux/amd64/main cmd/cli.go
	GOOS=darwin GOARCH=amd64 go build -o bin/darwin/amd64/main cmd/cli.go 
	GOOS=windows GOARCH=amd64 go build -o bin/windows/amd64/main cmd/cli.go 
run-cli:
	go run cmd/cli.go "select * from bingbong"
