package querybuilder

import (
	models "github.com/r00tk3y/prying-deep/pkg/querybuilder"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestQueryBuilderAddCriteria(t *testing.T) {
	assert := assert.New(t)
	qb := models.QueryBuilder{}

	emailCriteria := map[string]interface{}{"email": "example@com"}
	phoneNumberCriteria := map[string]interface{}{"international": "+8898442443"}

	criteriaMap := map[string]map[string]interface{}{
		"Email":       emailCriteria,
		"PhoneNumber": phoneNumberCriteria,
	}

	qb.AddCriteria(criteriaMap)

	assert.Equal(qb.PhoneNumberCriteria, phoneNumberCriteria)

	assert.Equal(qb.EmailCriteria, emailCriteria)
}

func TestQueryBuilderParseKeyValueCriteria(t *testing.T) {
	assert := assert.New(t)
	qb := models.QueryBuilder{}

	criteria := map[string]interface{}{"email": "test@email.com"}

	modelName := "Email"
	qb.ParseKeyValueCriteria(modelName, criteria)

	assert.Equal(qb.EmailCriteria, criteria)
}
