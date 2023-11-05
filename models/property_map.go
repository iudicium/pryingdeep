package models

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
)

// PropertyMap is the postgres implementation of jsonb in go.
type PropertyMap map[string]interface{}

func (p PropertyMap) Value() (driver.Value, error) {
	j, err := json.Marshal(p)
	return j, err
}

// Scan unmarshals a JSON-encoded byte slice into a PropertyMap.
// It assigns the result to the PropertyMap (p) or returns an error.
func (p *PropertyMap) Scan(src interface{}) error {
	source, ok := src.([]byte)
	if !ok {
		return errors.New("type assertion .([]byte) failed")
	}

	var i interface{}
	if err := json.Unmarshal(source, &i); err != nil {
		return err
	}

	*p, ok = i.(map[string]interface{})
	if !ok {
		return errors.New("type assertion .(map[string]interface{}) failed")
	}

	return nil
}
