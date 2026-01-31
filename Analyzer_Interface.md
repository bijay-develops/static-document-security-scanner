` internal/analyzer/analyzer.go `

### Design decision:

- Supports() allows extensibility
- Analyze() receives raw bytes
- Analyzer returns structured result

<i>We can plug in Excel, PowerPoint, ELF, anything later.</i>