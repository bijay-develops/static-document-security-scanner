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

