package dataclean

import (
	"encoding/csv"
	"fmt"
	"os"
	"strings"
)

// ReadCSV reads a CSV file and returns all rows as [][]string
func ReadCSV(filePath string) ([][]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to open CSV file: %w", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.TrimLeadingSpace = true
	reader.FieldsPerRecord = -1 // allow variable number of fields

	rawRecords, err := reader.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("failed to read CSV content: %w", err)
	}

	var cleanedRecords [][]string
	for _, row := range rawRecords {
		// skip empty lines
		if len(row) == 0 || (len(row) == 1 && strings.TrimSpace(row[0]) == "") {
			continue
		}
		cleaned := make([]string, len(row))
		for i, cell := range row {
			cleaned[i] = strings.TrimSpace(cell)
		}
		cleanedRecords = append(cleanedRecords, cleaned)
	}
	
	return cleanedRecords, nil
}
