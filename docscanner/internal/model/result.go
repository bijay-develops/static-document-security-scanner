package model

// ScanResult represents the result of a document scan.  -- Core Data Model
type ScanResult struct {
	FilePath string `json:"file_path"`
	FileType string `json:"file_type"`
	SHA256 string `json:"sha256"`
	Indicators []string `json:"indicators"`
	RiskScore int `json:"risk_score"`
}
