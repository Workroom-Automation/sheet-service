package sheet

import (
	"fmt"
	"github.com/pkg/errors"
)

type LessThanCondition struct {
	Value float64 `json:"value"`
}

// TODO make this whole config driven ....
var validNumberFieldTypeConditions = []ConditionType{
	LessThan,
	GreaterThan,
	LessThanAndEqualTo,
	GreaterThanAndEqualTo,
	EqualTo,
	NotEqualTo,
	Between,
	NotInBetween,
}

var validSelectionFieldTypeConditions = []ConditionType{
	Selected,
	NotSelected,
}

var validMultiSelectionFieldTypeConditions = []ConditionType{
	Selected,
	NotSelected,
}

var validDateSelectionFieldTypeConditions = []ConditionType{
	Before,
	After,
	Between,
	EqualTo,
}

func IsValidConditionForFieldType(fieldType FieldType, conditionType ConditionType) bool {
	switch fieldType {
	case Number:
		for _, validConditionType := range validNumberFieldTypeConditions {
			if validConditionType == conditionType {
				return true
			}
		}
	case Selection:
		for _, validConditionType := range validSelectionFieldTypeConditions {
			if validConditionType == conditionType {
				return true
			}
		}
	case MultiSelection:
		for _, validConditionType := range validMultiSelectionFieldTypeConditions {
			if validConditionType == conditionType {
				return true
			}
		}
	case DateSelection:
		for _, validConditionType := range validDateSelectionFieldTypeConditions {
			if validConditionType == conditionType {
				return true
			}
		}
	}
	return false
}

func ValidateTriggerConditionForFieldType(fieldType FieldType, conditionType ConditionType, conditionPayload interface{}) error {
	isValidConditionTypeForFieldType := IsValidConditionForFieldType(fieldType, conditionType)
	if !isValidConditionTypeForFieldType {
		return errors.New(fmt.Sprintf("invalid condition type %s, for field type %s", conditionType, fieldType))
	}
	switch conditionType {
	case LessThan:
		if _, ok := conditionPayload.(*SelectFieldFormData); ok {
			return nil
		}
	}
	return errors.New(fmt.Sprintf("invalid condition for field type %s, condition %v", fieldType, conditionPayload))
}
