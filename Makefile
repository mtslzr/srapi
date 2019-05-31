
GOBIN=go
PACKAGE=srapi

all: lint test vend build

build:
	${GOBIN} build -o ${PACKAGE} -v

deps: tidy vendor

lint:
	golangci-lint run

test:
	${GOBIN} test -v ./...

run: build
	./${PACKAGE}

tidy:
	${GOBIN} mod tidy -v

vend:
	${GOBIN} mod vendor -v