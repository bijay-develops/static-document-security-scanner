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

### Key files and symbols

- `cmd/scanner/main.go`
	- `func main()` – CLI entry, wiring of all components.

- `internal/scanner/walker.go`
	- `func WalkDirectory(root string, fileChan chan<- string) error` – recursive directory traversal.

- `internal/scanner/workerpool.go`
	- `func StartWorkerPool(numWorkers int, files <-chan string, analyzers []analyzer.Analyzer, results chan<- *model.ScanResult, wg *sync.WaitGroup)` – concurrent analysis workers.

- `internal/analyzer/analyzer.go`
	- `type Analyzer interface` – contract for all analyzers.

- `internal/analyzer/word.go`
	- `type WordAnalyzer struct{}` – OOXML / macro detection.

- `internal/analyzer/pdf.go`
	- `type PDFAnalyzer struct{}` – heuristic PDF keyword detection.

- `internal/model/result.go`
	- `type ScanResult struct` – core data model serialized as JSON.