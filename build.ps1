$env:GOOS="linux"
$env:GOARCH="amd64"
go build -o dist .
$env:GOOS="windows"
$env:GOARCH="amd64"
go build -o dist .