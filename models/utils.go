package models

import (
	"github.com/gofrs/uuid"
	"time"
)

// GenerateUuid is used for creating unique IDs to be used mainly in generated HTML elements
func GenerateUuid() (string, error) {
	curId, err := uuid.NewV4()
	if err != nil {
		return "", err
	}
	return curId.String(), nil
}

// ConvertToDateTime inputs a string formatted as "YYYY-MM-DDTHH:mm" and returns a parsed time.Time value
func ConvertToDateTime(dateString string) (time.Time, error) {
	layout := "2006-01-02T15:04"
	return time.Parse(layout, dateString)
}
