`docscanner/internal/scanner/walker.go`

### Responsibility

- Traverse a root directory recursively and emit every file path on a channel.

### Behavior

- Uses `filepath.Walk` starting from the provided `root`.
- For each visited path:
	- If there is an error, it is ignored and walking continues.
	- If the entry is not a directory, the full path is sent on `fileChan`.
- When walking completes, `fileChan` is closed.

The walker never reads file contents and has no knowledge of analyzers or risk scoring.