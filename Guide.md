# Static Document Security Scanner – User Guide

This guide explains how to install, run, and extend the static document security scanner.

## 1. Prerequisites

- OS: Linux, macOS, or Windows
- Go: version 1.18 or newer

Check your Go installation:

```bash
go version
```

## 2. Project Layout (Quick View)

From the repository root:

```text
.
├── docscanner/
│   ├── cmd/
│   │   └── scanner/
│   │       └── main.go         # CLI entrypoint
│   ├── internal/
│   │   ├── analyzer/           # Pluggable analyzers
│   │   │   ├── analyzer.go     # Analyzer interface
│   │   │   ├── pdf.go          # PDFAnalyzer
│   │   │   └── word.go         # WordAnalyzer
│   │   ├── scanner/
│   │   │   ├── walker.go       # Directory walker
│   │   │   └── workerpool.go   # Worker pool
│   │   └── model/
│   │       └── result.go       # ScanResult model
│   └── go.mod                  # Go module definition
├── samples/                    # Example documents to scan
└── Guide.md                    # This guide
```

## 3. One‑Time Setup

From the `docscanner` directory:

```bash
cd docscanner
# If go.mod does not exist yet, create it (run once):
go mod init docscanner

# Download and tidy dependencies:
go mod tidy
```

If `go.mod` already exists, you only need `go mod tidy`.

## 4. Where to Store Documents

You can scan **any directory** on your system. Two common options:

1. Use the provided `samples/` directory in the repo root:
   - Put `.pdf`, `.docx`, and `.docm` files in `samples/` (and its subfolders).
   - Example path:
     - `/home/<user>/Projects/static-document-security-scanner/samples`

2. Use your own folder anywhere, for example:
   - `~/Documents/docscanner-input`

The scanner just needs the directory path.

## 5. Running the Scanner

Always run the CLI from inside `docscanner/`.

### 5.1 Scan the example samples

From the repository root:

```bash
cd docscanner
go run ./cmd/scanner ../samples
```

### 5.2 Scan any other directory

Replace `<directory-to-scan>` with an absolute or relative path:

```bash
cd docscanner
go run ./cmd/scanner <directory-to-scan>
```

Examples:

```bash
# Scan the current directory
cd docscanner
go run ./cmd/scanner .

# Scan your Documents folder
cd docscanner
go run ./cmd/scanner ~/Documents
```

## 6. Saving Results to a File

By default, results are printed as JSON to the terminal (stdout).

To save them to a file:

```bash
cd docscanner
go run ./cmd/scanner ../samples > results.json
```

Open `results.json` in your editor to inspect the output.

## 7. Understanding the JSON Output

The scanner prints an array of `ScanResult` objects. Each object has:

```json
{
  "file_path": "../samples/Hotel_Management_Report.docx",
  "file_type": "word",
  "sha256": "<sha256 hash of the file>",
  "indicators": ["..."],
  "risk_score": 0
}
```

- `file_path` – Path of the scanned file.
- `file_type` – Logical type of the document (e.g., `word`, `pdf`).
- `sha256` – SHA‑256 hash of the file contents.
- `indicators` – List of strings describing what was found (e.g., suspicious features).
- `risk_score` – Simple numeric score based on indicators (higher means more suspicious).

## 8. What Is Detected Today?

### 8.1 Word documents (`.docx`, `.docm`)

Implemented in `internal/analyzer/word.go`:

- Detects presence of `vbaProject.bin` within the OOXML ZIP structure.
- If found, adds an indicator like `"Embedded VBA Macro (vbaProject.bin)"` and assigns a higher `risk_score`.

### 8.2 PDF documents (`.pdf`)

Implemented in `internal/analyzer/pdf.go`:

- Looks for suspicious markers such as:
  - `/JavaScript`, `/JS`, `/Launch`, `/OpenAction`, `/AA`, `/EmbeddedFile`.
- Each matched indicator increases `risk_score`.

## 9. Extending the Scanner

You can add support for new document types by implementing the `Analyzer` interface.

### 9.1 Analyzer interface

Defined in `internal/analyzer/analyzer.go`:

```go
type Analyzer interface {
    Supports(filename string) bool
    Analyze(filepath string, data []byte) (*model.ScanResult, error)
}
```

### 9.2 Steps to add a new analyzer

1. **Create a new analyzer file** in `internal/analyzer/`, for example `excel.go`.
2. **Implement** `Supports` to match your file extensions (e.g., `.xlsx`, `.xlsm`).
3. **Implement** `Analyze` to inspect `data` and return a `ScanResult`.
4. **Register the analyzer** in `cmd/scanner/main.go` by adding it to the `analyzers` slice:

```go
analyzers := []analyzer.Analyzer{
    &analyzer.WordAnalyzer{},
    &analyzer.PDFAnalyzer{},
    &analyzer.ExcelAnalyzer{}, // new
}
```

No changes are needed in the walker or worker pool – they automatically use the new analyzer.

## 10. Troubleshooting

- **Go command not found** – Install Go and ensure `go` is on your `PATH`.
- **Import path errors** – Make sure you ran `go mod init docscanner` once inside `docscanner/`.
- **Permission denied reading files** – Run the scanner on directories your user can read.
- **No results appear** – Verify the directory actually contains `.pdf`, `.docx`, or `.docm` files.

## 11. Next Steps

- Add more analyzers for other document types (e.g., Excel, PowerPoint).
- Enhance `risk_score` to consider file size, number of indicators, or custom rules.
- Integrate this CLI into a larger pipeline or CI step for automated document checks.
