## Static Document Security Scanner

This project is a small, composable static document scanner implemented in Go.

### What it does

- Walks a directory tree recursively
- Uses a worker pool to scan files concurrently based on `runtime.NumCPU()`
- Analyzes files using pluggable analyzers
- Computes SHA256 hashes for every scanned file
- Detects macro-enabled Word documents by locating `vbaProject.bin` inside the ZIP structure
- Detects suspicious PDF indicators using simple heuristic string matching
- Emits structured JSON results on stdout

### Quick usage

From the `docscanner` directory:

```bash
go run ./cmd/scanner <directory>
```

Example:

```bash
go run ./cmd/scanner ./samples
```

The CLI prints an array of JSON objects, each matching the `ScanResult` model defined in `internal/model/result.go`.

### Extensibility

- New document types can be added by implementing the `Analyzer` interface in `internal/analyzer/analyzer.go`.
- The worker pool and directory walker do not need to change when new analyzers are introduced.

This is intentionally a foundation: a minimal but solid base to grow more advanced detection logic.