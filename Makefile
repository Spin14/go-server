.PHONY: default fix lint test cover

default: fix lint test

fix:
	@ echo ">> fixing"
	@ gofmt -l -w .
	@ go mod tidy
	@ echo ">> done"

lint:
    # binary will be $(go env GOPATH)/bin/golangci-lint
    # curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.24.0
	@ echo ">> running linter"
	@ golangci-lint run
	@ echo ">> done"


test:
	@ echo ">> running tests"
	@ go test `go list ./... | grep -v internal`
	@ echo ">> done"

cover:
	@ echo ">> running tests and coverage"
	@ go test --count=1 --v --cover --coverprofile=cover.out `go list ./... | grep -v internal`
	@ go tool cover --html=cover.out
	@ echo ">> done"
