package analyzer

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"path/filepath"
	"strings"

	"docscanner/internal/model"
)

type PDFAnalyzer struct{}

var suspiciousIndicators = []string{
	"/JavaScript",
	"/JS",
	"/Launch",
	"/OpenAction",
	"/AA",
	"/EmbeddedFile",
}

func (p *PDFAnalyzer) Supports(fileName string) bool {
	return strings.ToLower(filepath.Ext(fileName)) == ".pdf"
}

func (p *PDFAnalyzer) Analyze(filePath string, data []byte) (*model.ScanResult, error) {
	indicators := []string{}
	for _, indicator := range suspiciousIndicators {
		if bytes.Contains(data, []byte(indicator)) {
			indicators = append(indicators, indicator)
		}
	}

	hash := sha256.Sum256(data)

	risk := len(indicators) * 15

	return &model.ScanResult{
		FilePath: filePath,
		FileType: "pdf",
		SHA256: hex.EncodeToString(hash[:]),
		Indicators: indicators,
		RiskScore: risk,
	}, nil 
}