env GOOS=target-OS GOARCH=target-architecture go build package-import-path
env GOOS=linux GOARCH=amd64 go build .
env GOOS=windows GOARCH=amd64 go build .
