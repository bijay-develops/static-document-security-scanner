`docscanner/cmd/scanner/main.go`

### Responsibilities

- Parse CLI arguments and validate that a root directory is provided
- Create channels for file paths and scan results
- Construct the list of active analyzers (Word + PDF)
- Kick off directory walking and the worker pool
- Wait for workers to finish and aggregate results
- Encode results as pretty-printed JSON to stdout

### High-level flow

1. Read `root := os.Args[1]`.
2. Create `files` (chan string) and `results` (chan *model.ScanResult).
3. Instantiate analyzers:
	- `&analyzer.WordAnalyzer{}`
	- `&analyzer.PDFAnalyzer{}`
4. Start directory walking with `scanner.WalkDirectory(root, files)`.
5. Start workers with `scanner.StartWorkerPool(runtime.NumCPU(), files, analyzers, results, &wg)`.
6. Wait for all workers via `WaitGroup`, then close `results`.
7. Collect all results and marshal them with `json.MarshalIndent`.

The CLI does not know about any analyzer internals; it only wires together the pieces.

