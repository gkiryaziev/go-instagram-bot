## Instagram Bot

---

[![Go Report Card](https://goreportcard.com/badge/github.com/gkiryaziev/go-instagram-bot)](https://goreportcard.com/report/github.com/gkiryaziev/go-instagram-bot)

This bot will check all recent activity in your Instagram accounts and record all data in the database.

### Installation:
```
go get github.com/gkiryaziev/go-instagram-bot
```

### Edit configuration:
```
Copy `config.default.yaml` to `config.yaml` and edit configuration.
```

### Build and Run:
```
go build && go-instagram-bot
```

### Packages:
You can use [glide](https://glide.sh/) packages manager to get all needed packages.
```
go get -u -v github.com/Masterminds/glide

cd go-instagram-bot && glide install
```

### Usage

#### 1. Drop tabbles
  `go-instagram-bot droptables`
#### 2. Auto-Migrate
  `go-instagram-bot migrate`
#### 3. Run fetcher
  `go-instagram-bot run`
