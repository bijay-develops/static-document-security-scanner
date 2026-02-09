# Supported Document Types

This project is a static document security scanner currently focused on the following document types.

## 1. Microsoft Word Documents

- Extensions: `.docx`, `.docm`
- Analyzer: `WordAnalyzer` (internal/analyzer/word.go)
- Detection logic:
  - Opens the OOXML container as a ZIP archive.
  - Searches for an embedded `vbaProject.bin` file.
  - If found, reports an indicator such as `"Embedded VBA Macro (vbaProject.bin)"` and assigns a higher risk score.

## 2. PDF Documents

- Extension: `.pdf`
- Analyzer: `PDFAnalyzer` (internal/analyzer/pdf.go)
- Detection logic:
  - Scans the raw PDF bytes for potentially dangerous features, for example:
    - `/JavaScript`
    - `/JS`
    - `/Launch`
    - `/OpenAction`
    - `/AA`
    - `/EmbeddedFile`
  - Each matched indicator increases the risk score for that file.

## 3. Other File Types

- The directory walker sees **all files**, but only files whose extensions are supported by an analyzer are actually scanned.
- Files that are not `.docx`, `.docm`, or `.pdf` are ignored by the current analyzers.

## 4. Extending Supported Types

To add support for new document types (for example, Excel or PowerPoint):

1. Implement the `Analyzer` interface in `internal/analyzer/` (see `analyzer.go`).
2. Decide which file extensions your analyzer supports in the `Supports` method.
3. Implement the `Analyze` method to inspect the file content and return a `ScanResult`.
4. Register the new analyzer in `cmd/scanner/main.go` by adding it to the `analyzers` slice.

This design lets you expand the set of supported document types without changing the directory walker or worker pool.