help:
	@echo 'make fmt                      Format everything'
	@echo 'make test                     Run tests (rebuilds mocks)'
	@echo 'make clean                    Format, mocks, vendoring'
	@echo 'make all                      Initialise, clean, test'

install-mockgen:
	go install go.uber.org/mock/mockgen@latest

.PHONY: vendor
deps:
	go mod tidy
	#go mod vendor

fmt:
	go fmt ./...

mocks:
	#./scripts/make-mocks.sh

clean: deps fmt mocks

.PHONY: test
test:
	go test ./...

all: clean test

docker-build:
	docker build -t goserver .

docker-run:
	docker run -p8080:8080 goserver
