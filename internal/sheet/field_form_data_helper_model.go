package sheet

import (
	"fmt"
	"github.com/pkg/errors"
)

type SelectFieldFormData struct {
	Options []interface{} `json:"options" binding:"required,min=1"`
}

type MultiSelectFieldFormData struct {
	Options []interface{} `json:"options" binding:"required,min=1"`
}

func ValidateFieldFormDataForFieldType(fieldType FieldType, formData interface{}) error {
	switch fieldType {
	case Selection:
		if _, ok := formData.(*SelectFieldFormData); ok {
			return nil
		}
	case Text:
	case Number:
	case DateSelection:
		// for text, number, date selection type field there is no helper data required...
		return nil
	case MultiSelection:
		if _, ok := formData.(*MultiSelectFieldFormData); ok {
			return nil
		}

	}
	return errors.New(fmt.Sprintf("invalid field form data for field type %s, field form options %v", fieldType, formData))
}
