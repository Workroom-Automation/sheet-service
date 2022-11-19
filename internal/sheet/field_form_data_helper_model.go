package sheet

import (
	"encoding/json"
	"fmt"
	"github.com/leapsquare/sheet-service/pkg/utils"
	"github.com/pkg/errors"
)

type SelectFieldFormData struct {
	Options []interface{} `json:"options" validate:"required,min=1"`
}

type MultiSelectFieldFormData struct {
	Options []interface{} `json:"options" validate:"required,min=1"`
}

func ValidateFieldFormDataForFieldType(fieldType FieldType, formData interface{}) error {
	jsonString, _ := json.Marshal(formData)
	validate := utils.GetJsonValidator()
	switch fieldType {
	case Selection:
		var selectStruct SelectFieldFormData
		err := json.Unmarshal(jsonString, &selectStruct)
		if err != nil {
			break
		}
		err = validate.Struct(selectStruct)
		if err == nil {
			return nil
		}
	case Text:
		return nil
	case Number:
		return nil
	case DateSelection:
		return nil
	case MultiSelection:
		var multiSelectStruct MultiSelectFieldFormData
		err := json.Unmarshal(jsonString, &multiSelectStruct)
		if err != nil {
			break
		}
		err = validate.Struct(multiSelectStruct)
		if err == nil {
			return nil
		}
	}
	return errors.New(fmt.Sprintf("invalid field form data for field type %s, field form options %v", fieldType, formData))
}
