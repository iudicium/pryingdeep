package querybuilder

import (
	"github.com/r00tk3y/prying-deep/models"
	"gorm.io/gorm"
	"reflect"
)

type QueryBuilder struct {
	MainModelCriteria map[string]interface{}
	SortBy            string
	SortOrder         string
	Limit             int
}

func (qb *QueryBuilder) ParseKeyValueCriteria(modelName string, criteria map[string]interface{}) {
	criteriaMapField := reflect.ValueOf(qb).Elem().FieldByName(modelName + "Criteria")

	if criteriaMapField.IsValid() {
		if criteriaMapField.IsNil() {
			criteriaMapField.Set(reflect.MakeMap(criteriaMapField.Type()))
		}

		for key, value := range criteria {
			criteriaMapField.SetMapIndex(reflect.ValueOf(key), reflect.ValueOf(value))
		}
	}
}
func (qb *QueryBuilder) AddCriteria(criteriaMap map[string]map[string]interface{}) {
	for modelName, criteria := range criteriaMap {
		switch modelName {
		case "Email":
			qb.ParseKeyValueCriteria(modelName, criteria)
		case "PhoneNumber":
			qb.ParseKeyValueCriteria(modelName, criteria)
		}
	}
}

func (qb *QueryBuilder) ConstructQuery(db *gorm.DB) *gorm.DB {
	query := db.Model(&models.WebPage{})

	if qb.SortOrder != "" || qb.SortBy != "" {
		query.Order(qb.SortOrder + "" + qb.SortBy)
	}

	if qb.Limit > 0 {
		query.Limit(qb.Limit)
	}

	return query
}
