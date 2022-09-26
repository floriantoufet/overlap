echo -e "Start tests"

echo -e "Start lint"
golangci-lint run

echo -e "Start units tests"
go test ./...

echo -e "Start integration tests"
go build -o test/godog/overlap main.go
cd test/godog
go mod download
godog run

echo -e "En tests"
