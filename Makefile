build:
	GOOS=linux GOARCH=amd64 go build -o bin/sqlparser-linux cmd/cli.go
	GOOS=darwin GOARCH=amd64 go build -o bin/sqlparser-mac cmd/cli.go 
	GOOS=windows GOARCH=amd64 go build -o bin/sqlparser-win.exe cmd/cli.go 

