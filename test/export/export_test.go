package export

import (
	"encoding/json"
	"fmt"
	"github.com/r00tk3y/prying-deep/configs"
	"github.com/r00tk3y/prying-deep/models"
	"github.com/r00tk3y/prying-deep/pkg/logger"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
	"os"
	"path/filepath"
	"testing"
)

func cleanUpDB(db *gorm.DB) {
	for i := 1; i < 100; i++ {
		query := fmt.Sprintf("DELETE FROM web_pages WHERE id = %d", i)
		db.Exec(query)
	}

}

func initDB() *gorm.DB {
	cfg := configs.GetConfig().DbConf
	dbURL := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.DbTestName)
	db := models.SetupDatabase(dbURL)

	db.Exec(fmt.Sprintf("ALTER SEQUENCE %s_id_seq RESTART WITH 1", "web_pages"))
	headers := `{"key1": "value1"}`

	for i := 1; i < 100; i++ {
		phoneNumbers := fmt.Sprintf("+123123%d", i)
		uniqueURL := fmt.Sprintf("http://test%d", i)
		uniqueTitle := fmt.Sprintf("test%d", i)
		uniquePgpKeys := fmt.Sprintf("{pgp_test%d}", i)
		uniqueEmails := fmt.Sprintf("{test%d@example.com}", i)
		uniqueWordpressHtml := fmt.Sprintf("{wordpress%d test}", i)

		db.Exec(`
            INSERT INTO web_pages (url, title, status_code, body, headers)
            VALUES ($1, $2, $3, $4, $5)
        `, uniqueURL, uniqueTitle, 200, "Venci vindi html", headers)

		db.Exec(`
            INSERT INTO phone_numbers (web_page_id, international_number, national_number, country_code)
            VALUES ($1, $2, $3, $4)
        `, i, phoneNumbers, phoneNumbers, "NL")

		db.Exec(`
            INSERT INTO cryptos (web_page_id, pgp_keys, certificates)
            VALUES ($1, $2, $3)
        `, i, uniquePgpKeys, "{test_certificate}")

		db.Exec(`
            INSERT INTO emails (web_page_id, emails)
            VALUES ($1, $2)
        `, i, uniqueEmails)

		db.Exec(`
            INSERT INTO wordpress_foot_prints (web_page_id, wordpress_html)
            VALUES ($1, $2)
        `, i, uniqueWordpressHtml)
	}

	return db
}
func TestMain(m *testing.M) {
	configs.SetupEnvironment()
	logger.InitLogger()
	defer logger.Logger.Sync()
	db := initDB()

	exitCode := m.Run()

	cleanUpDB(db)
	os.Exit(exitCode)
}

func TestJsonAndPreloadDBWithOneElement(t *testing.T) {
	tmpDir := t.TempDir()

	preloadedWebPage, err := models.PreloadWebPage(1)
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
