package main

import (
	"encoding/json"
	"fmt"
	"os"
	"runtime"
	"sync"

	"docscanner/internal/analyzer"
	"docscanner/internal/model"
	"docscanner/internal/scanner"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: scanner <directory>")
		os.Exit(1)
	}

	root := os.Args[1]

	files := make(chan string, 100)
	results := make(chan *model.ScanResult, 100)

	var wg sync.WaitGroup
	
	analyzers := []analyzer.Analyzer{
		&analyzer.WordAnalyzer{},
		&analyzer.PDFAnalyzer{},
	}

	go scanner.WalkDirectory(root, files)

	scanner.StartWorkerPool(runtime.NumCPU(), files, analyzers, results, &wg)

	go func() {
		wg.Wait()
		close(results)
	}()

	var allResults []*model.ScanResult

	for r := range results {
		allResults = append(allResults, r)
	}

	output, err := json.MarshalIndent(allResults, "", " ")
	if err != nil {
		fmt.Println("JSON error: ", err)
		return
	}

	fmt.Println(string(output))
}