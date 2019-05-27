# SRAPI

SRAPI is a Go-based API for the Sports Reference network of websites.

_Current Version: 0.2.0 (Pre-Release)_

| Master | Develop |
|:-:|:-:|
|[![Master](https://travis-ci.com/mtslzr/srapi.svg?branch=master)](https://travis-ci.com/mtslzr/srapi)|[![Develop](https://travis-ci.com/mtslzr/srapi.svg?branch=develop)](https://travis-ci.com/mtslzr/srapi)|

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
Starting server on localhost:5000...
```

### Executable

```bash
./srapi
Starting server on localhost:5000...
```

### Endpoints

#### Get Standings

Returns current standings. Requires two-letter ID for sport (e.g. `bs` for baseball).

```
GET /{sport}/standings
```

#### Get Teams

Returns all current teams. Requires two-letter ID for sport (e.g. `bs` for baseball).

```
GET /{sport}/teams
```

#### Get Years

Returns all available years. Requires two-letter ID for sport (e.g. `bs` for baseball).

```
GET /{sport}/years
```

## Contribution

### Testing

```bash
go test -v ./...
```

### Deployment

_Coming soon..._