echo -e "Start tests"

echo -e "Start lint"
golangci-lint run

echo -e "Start units tests"
go test ./...
