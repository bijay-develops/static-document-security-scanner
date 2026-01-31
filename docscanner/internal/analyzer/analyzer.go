// Analyzer interface
package analyzer

import "docscanner/internal/model"

type Analyzer interface {
	Supports(filename string) bool			// Check if the analyzer supports the given file type
	Analyze(filepath string, data []byte) (*model.ScanResult, error)   // Analyze the file and return a ScanResult
}



//Supports() allows extensibility
// Analyze() receives raw bytes
// Analyzer returns structured result