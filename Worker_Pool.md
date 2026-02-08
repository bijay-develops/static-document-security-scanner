`docscanner/internal/scanner/workerpool.go`

### Design reasoning

- Workers are isolated: each goroutine reads file paths from the `files` channel and processes them independently.
- Analyzer logic is injected: the worker pool only depends on the `Analyzer` interface.
- Extensible: adding a new analyzer is just appending to the `analyzers` slice in `main.go`.
- Proper synchronization: a `sync.WaitGroup` tracks worker lifetimes; results are written to a shared `results` channel.

### Behavior

- Spawn `numWorkers` goroutines.
- For each incoming file path:
	- Read file contents with `os.ReadFile`.
	- Find the first analyzer whose `Supports` method returns true.
	- Call `Analyze`; if it returns a non-nil result and no error, send it to `results`.

The worker pool does not know about directory structure or JSON output; it focuses purely on concurrent analysis.

