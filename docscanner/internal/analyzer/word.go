package analyzer

import (
	"archive/zip"
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"path/filepath"
	"strings"

	"docscanner/internal/model"
)

type WordAnalyzer struct {}

func (w *WordAnalyzer) Supports(fileName string) bool {
	ext := strings.ToLower(filepath.Ext(fileName))
	return ext == ".docx" || ext == ".docm"
}

func (w *WordAnalyzer) Analyze(filePath string, data []byte) (*model.ScanResult, error) {
	readerAt := bytes.NewReader(data)
	zr, err := zip.NewReader(readerAt, int64(len(data)))
	if err != nil {
		return nil, err
	}

	indicators := []string{}
	for _, file := range zr.File {
		if strings.Contains(strings.ToLower(f.Name), "vbaproject.bin") {
			indicators = append(indicators, "Embedded VBA MAcro (vbaProject.bin)")
			break
		}
	}

	hash := sha256.Sum256(data)

	risk := 0

	if len(indicators) > 0 {
		risk = 70
	}

	return &model.ScanResult{
		FilePath: filePath,
		FileType: "word",
		SHA256: hex.EncodeToString(hash[:]),
		Indicators: indicators,
		RiskScore: risk,
	}, nil
}

