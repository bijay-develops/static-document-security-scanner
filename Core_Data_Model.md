`docscanner/internal/model/result.go`

### ScanResult

```go
type ScanResult struct {
	FilePath   string   `json:"file_path"`
	FileType   string   `json:"file_type"`
	SHA256     string   `json:"sha256"`
	Indicators []string `json:"indicators"`
	RiskScore  int      `json:"risk_score"`
}
```

### Semantics

- `FilePath` – absolute or relative path to the analyzed file.
- `FileType` – logical type label (e.g. "word", "pdf").
- `SHA256` – hex-encoded SHA256 hash of the full file contents.
- `Indicators` – human-readable strings describing findings (e.g. suspicious keywords or macro presence).
- `RiskScore` – coarse numeric score computed by each analyzer.

The CLI aggregates a slice of `ScanResult` values and renders them as JSON.