package tests

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"

	"github.com/iudicium/pryingdeep/configs"
	"github.com/iudicium/pryingdeep/pkg/exporters"
	"github.com/iudicium/pryingdeep/pkg/logger"
	"github.com/iudicium/pryingdeep/pkg/querybuilder"
	"github.com/iudicium/pryingdeep/tests/test_helpers"
)

var db *gorm.DB

func parseJsonFile(filepath string) ([]map[string]interface{}, error) {
	var jsonData []map[string]interface{}
	jsonFile, err := os.ReadFile(filepath)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(jsonFile, &jsonData); err != nil {
		return nil, err
	}
	return jsonData, nil

}

func assertJSONStructure(data map[string]interface{}, assert *assert.Assertions) {
	expectedKeys := []string{"pageData", "crypto", "email", "phoneNumbers", "wordpress"}
	for _, key := range expectedKeys {
		assert.Contains(data, key, "Expected key '%s' not found", key)
	}
}

func TestMain(m *testing.M) {
	configs.SetupEnvironment()
	logger.InitLogger(false)
	defer logger.Logger.Sync()
	db = test_helpers.InitDB()

	exitCode := m.Run()

	test_helpers.CleanUpDB(db)
	os.Exit(exitCode)
}

func TestJsonAndPreloadDBWithOneElement(t *testing.T) {
	tmpDir := t.TempDir()

	preloadedWebPage, err := test_helpers.PreloadWebPage(db, 1)
	if err != nil {
		t.Fatal(err)
	}
	preloadedJSON, err := json.MarshalIndent(preloadedWebPage, "", " ")
	if err != nil {
		t.Fatal(err)
	}
	path := filepath.Join(tmpDir, "test.json")
	err = os.WriteFile(path, preloadedJSON, 0644)
	if err != nil {
		t.Fatal(err)
	}
	content, err := os.ReadFile(path)
	if err != nil {
		t.Fatal()
	}
	var result map[string]interface{}
	if err := json.Unmarshal(content, &result); err != nil {
		t.Fatal(err)
	}
	assert.Contains(t, result, "crypto")
	assert.Contains(t, result, "email")
	assert.Contains(t, result, "phoneNumbers")
	assert.Contains(t, result, "wordpress")
	assert.Contains(t, result, "email")
}

func TestConvertQueryBuilderDataToJson(t *testing.T) {
	testCases := []struct {
		Name            string
		WebPageCriteria map[string]interface{}
		SortBy          string
		SortOrder       string
		Limit           int
		ItemLength      int
		Associations    string
		Offset          int
	}{
		{
			Name: "JsonTestWithOneLimit",
			WebPageCriteria: map[string]interface{}{
				"URL": "LIKE test",
			},
			SortBy:       "url",
			SortOrder:    "",
			Limit:        1,
			ItemLength:   1,
			Associations: "all",
			Offset:       0,
		},
		{
			Name: "JsonTestWithUnlimited",
			WebPageCriteria: map[string]interface{}{
				"URL": "LIKE test",
			},
			SortBy:       "url",
			SortOrder:    "",
			Limit:        0,
			ItemLength:   99,
			Associations: "all",
			Offset:       0,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			//Uncomment to analyze the actual json
			//tmpDir := ""
			tmpDir := t.TempDir()
			tmpPath := filepath.Join(tmpDir, "test.json")
			assert := assert.New(t)

			qb := querybuilder.NewQueryBuilder(
				tc.WebPageCriteria,
				tc.Associations,
				tc.SortBy,
				tc.SortOrder,
				tc.Limit,
				tc.Offset,
			)

			result := qb.ConstructQuery(db)
			exporter := exporters.NewExporter(tmpPath)
			err := exporter.ToJSON(result)
			if err != nil {
				t.Fatal(err)
			}

			data, err := parseJsonFile(tmpPath)
			if err != nil {
				t.Fatal(err)
			}
			assert.Equal(len(data), tc.ItemLength)
			for _, item := range data {
				assertJSONStructure(item, assert)
			}

		})
	}
}

//func TestConvertQueryBuilderToCSV(t *testing.T) {
//
//}
