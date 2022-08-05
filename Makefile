GOPATH:=$(shell go env GOPATH)
APP?=urlshortener

.PHONY: init
## init: initialize the application(urlshortener)
init:
	go mod download

.PHONY: build
## build: build the application(urlshortener)
build:
	go build -o build/${APP} cmd/main.go

.PHONY: run
## run: run the application(urlshortener)
run:
	go run -v -race cmd/main.go

.PHONY: format
## format: format files
format:
	@go install golang.org/x/tools/cmd/goimports@latest
	@go install github.com/banksalad/go-banksalad/cmd/importsort@latest
	goimports -local github.com/banksalad -w .
	importsort -s github.com/banksalad -w $$(find . -name "*.go")
	gofmt -s -w .
	go mod tidy

.PHONY: test
## test: run tests
test:
	@go install github.com/rakyll/gotest@latest
	gotest -race -cover -v ./...

.PHONY: coverage
## coverage: run tests with coverage
coverage:
	@go install github.com/rakyll/gotest@latest
	gotest -race -coverprofile=coverage.txt -covermode=atomic -v ./...

.PHONY: lint
## lint: check everything's okay
lint:
	golangci-lint run ./...
	go mod verify

.PHONY: generate
## generate: generate source code for mocking
generate:
	@go install golang.org/x/tools/cmd/stringer@latest
	@go install github.com/golang/mock/mockgen@latest
	go generate ./...
	@$(MAKE) format

.PHONY: generate-db-models
## generate-db-models: generate source code for DB models
generate-db-models:
	@go install github.com/volatiletech/sqlboiler/v4@v4.8.6
	@go install github.com/volatiletech/sqlboiler/v4/drivers/sqlboiler-mysql@v4.8.6
	cat sqlboiler.toml.tmpl | \
		docker run --rm -i \
			-e DB_NAME \
			-e DB_HOST \
			-e DB_PORT \
			-e DB_USER \
			-e DB_PASSWORD \
			-e DB_SSL_MODE \
			hairyhenderson/gomplate:v3.6.0-slim -f - > sqlboiler.toml
	sqlboiler --templates ./server/db/templates \
		--templates $$(go env GOPATH)/pkg/mod/github.com/volatiletech/sqlboiler/v4@v4.8.6/templates/main \
		--wipe --no-tests --no-auto-timestamps -p mysql -o ./server/db/mysql mysql

.PHONY: help
## help: prints this help message
help:
	@echo "Usage: \n"
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':'

.PHONY: deploy-to-k8s
## deploy-to-k8s: deploy application to kubernetes
deploy-to-k8s:
	cat k8s.yaml.tmpl | gomplate -f - | kubectl apply -f -
