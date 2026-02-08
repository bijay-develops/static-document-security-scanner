`docscanner/internal/analyzer/pdf.go`

### What we detect

- Suspicious keywords inside the raw PDF byte stream.

### Heuristic indicators

- `/JavaScript`
- `/JS`
- `/Launch`
- `/OpenAction`
- `/AA`
- `/EmbeddedFile`

The analyzer:

- Checks file extension is `.pdf` via `Supports`.
- Scans the raw bytes for each of the indicators above.
- Computes the file's SHA256 hash.
- Sets `RiskScore` to `15 * len(indicators)`.

This is intentionally heuristic, not signature-based: static, cheap, and fast.