package aggregates

import "github.com/google/uuid"

type CustomField struct {
    ID          string
    Name        string
    Type        string
    Required    bool
}

func CreateNewCustomField(name, fieldType string, required bool) *CustomField {
    return &CustomField{
        ID:       uuid.New().String(),
        Name:     name,
        Type:     fieldType,
        Required: required,
    }
}
