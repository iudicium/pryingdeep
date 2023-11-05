package querybuilder

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"github.com/pryingbytez/pryingdeep/models"
	"github.com/pryingbytez/pryingdeep/pkg/logger"
)

// QueryBuilder  stores SQL parameters that are used for performing gorm SQL statements.
type QueryBuilder struct {
	//WebPageCriteria is a map that accept different fields.
	//It takes in key value pairs. You can also specify the LIKE keyword like this:
	//title: LIKE example
	WebPageCriteria map[string]interface{}
	//Associations - pryingtools shortened
	Associations string

	SortBy    string
	SortOrder string
	Limit     int
}

func NewQueryBuilder(webPageCriteria map[string]interface{}, a, sortBy, sortOrder string, limit int) *QueryBuilder {
	return &QueryBuilder{
		WebPageCriteria: webPageCriteria,
		Associations:    a,
		SortBy:          sortBy,
		SortOrder:       sortOrder,
		Limit:           limit,
	}
}

// ConstructQuery (ConstructQuery) Constructs the queries based on the fields
func (qb *QueryBuilder) ConstructQuery(db *gorm.DB) []models.WebPage {
	var pages []models.WebPage
	var err error
	query := db.Model(&models.WebPage{})
	query, err = ParseAndPreloadAssociations(db, qb.Associations)
	if err != nil {
		logger.Errorf("err during preloading associations: %s", err)
	}
	if qb.WebPageCriteria != nil {
		for key, value := range qb.WebPageCriteria {
			condition := BuildCondition(db, key, value)
			query.Where(condition)

		}
	}
	if qb.SortBy != "" {
		if qb.SortOrder != "" {
			query.Order(qb.SortBy + " " + qb.SortOrder)
		} else {
			query.Order(qb.SortBy)
		}

	}

	if qb.Limit > 0 {
		query.Limit(qb.Limit)
	}

	query.Find(&pages)
	return pages

}

// Raw is a helper for executing raw queries inside the database. You can define
// Your query anywhere you want and call this method to execute custom queries
// Note: This will not provide structured keys like ConstructQuery.
// However, this function does give you more control on what fields you can choose from other models and export them later on.
// This function, also does not support INSERT statements.
func (qb *QueryBuilder) Raw(db *gorm.DB, relativePath string) (error, []map[string]interface{}) {

	results := make([]map[string]interface{}, 0)
	path, err := filepath.Abs(relativePath)
	if err != nil {
		return err, results
	}
	queryBytes, err := os.ReadFile(path)
	if err != nil {
		return err, results
	}
	query := string(queryBytes)

	err = db.Raw(query).Scan(&results).Error
	if err != nil {
		return err, results
	}

	return nil, results
}

// BuildCondition is the builder for the provided WebPageCriteria, it also supports the usage
// of LIKE statements without needing the extra %
func BuildCondition(query *gorm.DB, field string, criteria interface{}) *gorm.DB {
	var pattern string

	if strValue, ok := criteria.(string); ok {
		if strings.Contains(strValue, "LIKE") {
			pattern = strings.TrimSpace(strings.TrimPrefix(strValue, "LIKE"))
			return query.Where(fmt.Sprintf("%s LIKE ?", field), "%"+pattern+"%")
		} else {
			return query.Where(fmt.Sprintf("%s = ?", field), strValue)
		}
	}

	return query.Where(fmt.Sprintf("%s = ?", field), criteria)
}

// ParseAndPreloadAssociations  is there as a helper to preload specific tables that the user
// Would like to be exported. If the "all" parameter is specified, it will exporters all the models
func ParseAndPreloadAssociations(query *gorm.DB, associations string) (*gorm.DB, error) {
	if associations == "" {
		associations = "ALL"
	}
	associations = strings.ToUpper(associations)
	if associations == "ALL" {
		return query.Preload(clause.Associations), nil
	}
	associationMap := map[string]string{
		"E":  "Emails",
		"WP": "Wordpress",
		"P":  "PhoneNumbers",
		"C":  "Crypto",
	}

	associationList := strings.Split(associations, ",")
	for _, assoc := range associationList {
		preload, found := associationMap[assoc]
		if found {
			query = query.Preload(preload)
		} else {
			return nil, fmt.Errorf("invalid association: %s", assoc)
		}
	}
	return query, nil
}
