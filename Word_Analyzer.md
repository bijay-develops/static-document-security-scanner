`docscanner/internal/analyzer/word.go`

### What we detect

- Presence of `vbaProject.bin` inside the OOXML ZIP structure for Word documents.

### How it works

- `Supports` returns true for `.docx` and `.docm` files.
- The file bytes are treated as a ZIP archive using `archive/zip`.
- All entries are scanned; if any name contains `vbaproject.bin` (case-insensitive), the analyzer records an indicator:
	- "Embedded VBA Macro (vbaProject.bin)".
- A SHA256 hash of the entire document is computed.
- If at least one indicator is present, `RiskScore` is set to `70`; otherwise `0`.

This is a coarse but useful signal for macro-enabled Office documents.