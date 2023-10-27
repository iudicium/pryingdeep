package testdb

import (
	"fmt"

	"github.com/pryingbytez/prying-deep/configs"
	"github.com/pryingbytez/prying-deep/models"
	"gorm.io/gorm"
)

// InitDB is  only meant for testing
func InitDB() *gorm.DB {
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

func CleanUpDB(db *gorm.DB) {
	for i := 1; i < 100; i++ {
		query := fmt.Sprintf("DELETE FROM web_pages WHERE id = %d", i)
		db.Exec(query)
	}

}
