`docscanner/internal/analyzer/analyzer.go`

### Interface

```go
type Analyzer interface {
	Supports(filename string) bool
	Analyze(filepath string, data []byte) (*model.ScanResult, error)
}
```

### Design decisions

- `Supports` decides, based on filename/extension, whether this analyzer owns the file.
- `Analyze` receives the full file contents as a `[]byte` plus the path.
- Each analyzer returns a `*model.ScanResult` describing its findings.

### Extensibility

- New analyzers live in `internal/analyzer/` and implement this interface.
- The worker pool only depends on the interface, so adding new analyzers does not change concurrency code.

<i>We can plug in Excel, PowerPoint, archives, binaries—anything—behind this interface.</i>