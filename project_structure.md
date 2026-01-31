```
docscanner/
│
├── cmd/
│   └── scanner/
│       └── main.go
│
├── internal/
│   ├── analyzer/
│   │   ├── analyzer.go
│   │   ├── word.go
│   │   └── pdf.go
│   │
│   ├── scanner/
│   │   ├── walker.go
│   │   └── workerpool.go
│   │
│   └── model/
│       └── result.go
│
├── go.mod

```

Why this layout?

- `cmd/` → entrypoints
- `internal/` → private business logic
- `analyzer/` → pluggable detection modules
- `scanner/` → orchestration + concurrency
- `model/` → shared data types

This is idiomatic Go service layout.