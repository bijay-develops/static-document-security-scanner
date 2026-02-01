package scanner

import (
	"os"
	"path/filepath"
)

func WalkDirectory(root string, fileChan chan<- string) error {
	defer close(fileChan)

	return filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return nil
		}
		if !info.IsDir() {
			fileChan <- path
		}
		return nil
	})
}
