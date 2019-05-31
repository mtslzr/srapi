
GOCMD=go
GOBUILD=${GOCMD} build
GOLINT=${GOCMD} lint
GOLINTCI=golangci-lint
GOMOD=${GOCMD} mod
GOTEST=${GOCMD} test

PACKAGE=srapi

DB=${PACKAGE}.db
DBCMD=sqlite3
SQL=${PACKAGE}.sql

all: lint test build

build:
	${GOBUILD} -o ${PACKAGE} -v

db:
	cat ${SQL} | ${DBCMD} ${DB}

deps: tidy vendor

dump:
	${DBCMD} ${DB} .dump > ${SQL}

lint:
	${GOLINT} *.go
	${GOLINTCI} run

prep: deps dump

test:
	${GOTEST} -v ./...

run: build
	./${PACKAGE}

tidy:
	${GOMOD} tidy -v

vend:
	${GOMOD} vendor -v