# SRAPI

SRAPI is a Go-based API for the Sports Reference network of websites.

_Current Version: 0.1.0 (Pre-Release)_

## Installation

### Clone Repo

```bash
git clone git@github.com:mtslzr/srapi
cd srapi
```

### Dependencies

_NOTE: Database filename is currently hard-coded. This will change in the future._

```bash
cat srapi.sql | sqlite3 srapi.db
go get -v ./...
```

### Build

```bash
go build -o srapi
```

## Usage

### Run

```bash
go run -v ./...
```

### Executable

```bash
./srapi
Starting server on localhost:5000...
```

### Endpoints

#### Dummy

```
GET - /
GET - /{sport}
```

## Contribution

### Testing

_NOTE: Not implemented yet._

```bash
go test -v ./...
```

### Deployment

_Coming soon..._