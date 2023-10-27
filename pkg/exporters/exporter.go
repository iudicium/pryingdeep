package exporters

import (
	"encoding/json"
	"os"

	"github.com/pryingbytez/prying-deep/models"
	"github.com/pryingbytez/prying-deep/pkg/logger"
)

// Exporter is a struct for exporting data from a database in a convenient format.
// It provides methods to export data to various formats such as JSON, CSV, etc.
type Exporter struct {
	FilePath string
}

func NewExporter(path string) *Exporter {
	return &Exporter{
		FilePath: path,
	}
}
func (e Exporter) ToJSON(data []models.WebPage) error {
	preloadedJSON, err := json.MarshalIndent(data, "", " ")
	if err != nil {
		return err
	}

	err = os.WriteFile(e.FilePath, preloadedJSON, 0644)
	if err != nil {
		return err
	}

	logger.Infof("File saved to: %s successfully!", e.FilePath)
	return nil
}

//TODO: fix this funciton, later on
//func (e Exporter) ToCSV(data []models.WebPage) error {
//	headers := []string{"pageData", "email", "crypto", "wordpress", "phone"}
//
//	outputFile, err := os.Create(e.FilePath)
//	if err != nil {
//		return err
//	}
//	defer outputFile.Close()
//
//	jsonData, err := json.Marshal(data)
//	if err != nil {
//		return err
//	}
//	writer := csv.NewWriter(outputFile)
//	defer writer.Flush()
//	if err = writer.Write(headers); err != nil {
//		return err
//	}
//	for key, value := range jsonData {
//		var csvRow []string
//		csvRow = append(csvRow, d.PageData, d.Emails)
//
//	}
//	return nil
//
//}
