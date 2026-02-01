### What we detect:
- Suspicious keywords inside raw PDF stream.

### Heuristic indicators:

- /JavaScript
- /JS
- /Launch
- /OpenAction
- /AA
- /EmbeddedFile

` internal/analyzer/pdf.go `

### This is intentionally heuristic, not signature-based. Static, cheap, fast.