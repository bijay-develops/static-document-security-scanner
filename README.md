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

### Code structure overview

High-level flow:

```text
main.go
	├─ parses CLI args
	├─ creates channels (files, results)
	├─ builds analyzers []Analyzer
	├─ starts directory walker (WalkDirectory)
	├─ starts worker pool (StartWorkerPool)
	└─ aggregates []ScanResult and prints JSON

WalkDirectory (internal/scanner/walker.go)
	└─ walks the filesystem and pushes file paths into files chan

StartWorkerPool (internal/scanner/workerpool.go)
	└─ spins up N workers
			 └─ for each file:
						├─ os.ReadFile
						├─ pick matching Analyzer via Supports
						└─ Analyzer.Analyze → *ScanResult → results chan

Analyzers (internal/analyzer/*.go)
	├─ WordAnalyzer  – detects vbaProject.bin in .docx/.docm
	└─ PDFAnalyzer   – scans for suspicious PDF keywords

Model (internal/model/result.go)
	└─ ScanResult – structure that is serialized to JSON
```

### Data flow (Mermaid)

```mermaid
flowchart TD
	CLI["CLI (main.go)"] --> ARGS["Parse args (root dir)"]
	ARGS --> CHANNELS["Create channels: files, results"]
	CHANNELS --> ANALYZERS["Build analyzers []Analyzer"]
	CHANNELS --> WALKER["WalkDirectory (walker.go)"]
	WALKER --> FILES["files chan <- file paths"]
	FILES --> WORKERPOOL["StartWorkerPool (workerpool.go)"]
	WORKERPOOL --> WORKER["N workers"]
	WORKER --> READ["os.ReadFile(file)"]
	READ --> DISPATCH{"Supports(file)?"}
	DISPATCH -->|yes| ANALYZE["Analyzer.Analyze(file, data)"]
	ANALYZE --> RESULTS["results chan <- *ScanResult"]
	RESULTS --> AGG["Aggregate []ScanResult"]
	AGG --> JSON["Marshal JSON & print"]
```