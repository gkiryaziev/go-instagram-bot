## Instagram News Bot

---

This bot will check all news in your Instagram accounts and record all activity in the database.

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

### Usage

#### 1. Drop tabbles
  `go-instagram-bot droptables`
#### 2. Auto-Migrate
  `go-instagram-bot migrate`
#### 3. Run fetcher
  `go-instagram-bot run`
