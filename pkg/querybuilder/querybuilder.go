package querybuilder

import (
	"fmt"
	"strings"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"github.com/pryingbytez/prying-deep/models"
	"github.com/pryingbytez/prying-deep/pkg/logger"
)

type QueryBuilder struct {
	WebPageCriteria map[string]interface{}
	Associations    string
	SortBy          string
	SortOrder       string
	Limit           int
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
	if qb.SortOrder != "" || qb.SortBy != "" {
		query.Order(qb.SortBy + "" + qb.SortOrder)
	}

	if qb.Limit > 0 {
		query.Limit(qb.Limit)
	}

	query.Find(&pages)
	return pages

}

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
