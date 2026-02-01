package scanner

import (
	"os"
	"sync"

	"docscanner/internal/analyzer"
	"docscanner/internal/model"
)

func StartWorkerPool(
	numWorkers int,
	files <-chan string,
	analyzers []analyzer.Analyzer,
	results chan<- *model.ScanResult,
	wg *sync.WaitGroup,
) {
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for file := range files {
				data, err := os.ReadFile(file)
				if err != nil {
					continue
				}

				for _, a := range analyzers {
					if a.Supports(file) {
						res, err := a.Analyze(file, data)
						if err == nil && res != nil {
							results <- res
						}
						break
					}
				}
			}
		}()
	}
}
