package tests

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"

	"github.com/r00tk3y/prying-deep/configs"
	"github.com/r00tk3y/prying-deep/internal/testdb"
	"github.com/r00tk3y/prying-deep/models"
	"github.com/r00tk3y/prying-deep/pkg/logger"
	"github.com/r00tk3y/prying-deep/pkg/querybuilder"
)

var db *gorm.DB

func constructQueryHelper(associations string, qb *querybuilder.QueryBuilder) []models.WebPage {
	pages := qb.ConstructQuery(db, associations)

	queryJSON, err := json.Marshal(pages)
	fmt.Println(string(queryJSON))
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(queryJSON, &pages)
	if err != nil {
		panic(err)
	}

	return pages
}

func TestMain(m *testing.M) {
	configs.SetupEnvironment()
	logger.InitLogger()
	defer logger.Logger.Sync()

	cfg := configs.GetConfig().DbConf
	dbURL := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.DbTestName)
	db = models.SetupDatabase(dbURL)

	db = testdb.InitDB()

	exitCode := m.Run()

	testdb.CleanUpDB(db)
	os.Exit(exitCode)
}

func TestQueryBuilder_ConstructQuery(t *testing.T) {
	assert := assert.New(t)
	t.Run("AllAssociations", func(t *testing.T) {
		associations := "all"
		qb := querybuilder.NewQueryBuilder(
			map[string]interface{}{
				"URL": "LIKE test",
			},
			"url",
			"",
			1,
		)

		result := constructQueryHelper(associations, qb)
		testEmail := result[0].Emails.Emails[0]
		testPGPKey := result[0].Crypto.PGPKeys[0]
		testWordpress := result[0].Wordpress.WordpressHtml[0]
		testPhoneNumber := result[0].PhoneNumbers.InternationalNumber
		assert.Equal(len(result), 1)
		assert.Equal(result[0].URL, "http://test1")
		assert.Equal(testEmail, "test1@example.com")
		assert.Equal(testPGPKey, "pgp_test1")
		assert.Equal(testWordpress, "wordpress1 test")
		assert.Equal(testPhoneNumber, "+1231231")

	})

	// Test case 2: Specific associations,
	//this test also looks for the Limit variable to be working
	t.Run("EmailAndWordPressAssociations", func(t *testing.T) {
		associations := "E,WP"
		qb := &querybuilder.QueryBuilder{
			WebPageCriteria: map[string]interface{}{
				"URL": "LIKE test",
			},
			SortBy: "url",
			Limit:  10,
		}
		result := constructQueryHelper(associations, qb)
		assert.Equal(len(result), 10)

		//This just means that its empty
		expectedCrypto := (*models.Crypto)(nil)
		expectedPhoneNumbers := (*models.PhoneNumber)(nil)

		assert.Equal(expectedCrypto, result[0].Crypto)
		assert.Equal(expectedPhoneNumbers, result[0].PhoneNumbers)
	})

	t.Run("CryptoAndPhoneNumbersAssociations", func(t *testing.T) {
		associations := "P,C"
		qb := &querybuilder.QueryBuilder{
			WebPageCriteria: map[string]interface{}{
				"title": "LIKE test ",
			},
			SortBy: "url",
			Limit:  5,
		}
		result := constructQueryHelper(associations, qb)

		assert.Equal(len(result), 5)

		//This just means that its empty
		expectedWordpress := (*models.WordpressFootPrint)(nil)
		expectedEmail := (*models.Email)(nil)
		assert.Equal(expectedWordpress, result[0].Wordpress)
		assert.Equal(expectedEmail, result[0].Emails)

	})
	t.Run("NoAssociationProvided", func(t *testing.T) {
		associations := ""
		qb := &querybuilder.QueryBuilder{
			WebPageCriteria: map[string]interface{}{
				"title": "LIKE test ",
			},
			SortBy: "url",
			Limit:  5,
		}
		result := constructQueryHelper(associations, qb)

		assert.Equal(len(result), 5)
		assert.Equal(result[0].URL, "http://test1")

	})
	//No criteria will just mean that the queryBuilder will fetch every object in the database
	t.Run("No criteria provided", func(t *testing.T) {
		associations := ""
		qb := &querybuilder.QueryBuilder{}
		result := constructQueryHelper(associations, qb)

		assert.Equal(len(result), 99)
		assert.Equal(result[0].URL, "http://test1")

	})
}
