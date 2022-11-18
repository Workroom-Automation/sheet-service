package sheet

import (
	"fmt"
	"github.com/pkg/errors"
	"time"
)

type LessThanCondition struct {
	Value float64 `json:"value"`
}

type GreaterThanCondition struct {
	Value float64 `json:"value"`
}

type LessThanAndEqualToCondition struct {
	Value float64 `json:"value"`
}

type GreaterThanAndEqualToCondition struct {
	Value float64 `json:"value"`
}

type EqualToCondition struct {
	Value float64 `json:"value"`
}

type NotEqualToCondition struct {
	Value float64 `json:"value"`
}

type BetweenCondition struct {
	UpperLimit float64 `json:"upper_limit"`
	LowerLimit float64 `json:"lower_limit"`
}

type NotInBetweenCondition struct {
	UpperLimit float64 `json:"upper_limit"`
	LowerLimit float64 `json:"lower_limit"`
}

type SelectedCondition struct {
	// would store the options in terms of the alphabet or numbers...
	Options []string `json:"options"`
}

type NotSelectedCondition struct {
	Options []string `json:"options"`
}

type BeforeCondition struct {
	DateTime time.Time `json:"date_time"`
}

type AfterCondition struct {
	DateTime time.Time `json:"date_time"`
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
		if _, ok := conditionPayload.(*LessThanCondition); ok {
			return nil
		}
	case GreaterThan:
		if _, ok := conditionPayload.(*GreaterThanCondition); ok {
			return nil
		}
	case LessThanAndEqualTo:
		if _, ok := conditionPayload.(*LessThanAndEqualToCondition); ok {
			return nil
		}
	case GreaterThanAndEqualTo:
		if _, ok := conditionPayload.(*GreaterThanAndEqualToCondition); ok {
			return nil
		}
	case EqualTo:
		if _, ok := conditionPayload.(*EqualToCondition); ok {
			return nil
		}
	case NotEqualTo:
		if _, ok := conditionPayload.(*NotEqualToCondition); ok {
			return nil
		}
	case Between:
		if _, ok := conditionPayload.(*BetweenCondition); ok {
			return nil
		}
	case NotInBetween:
		if _, ok := conditionPayload.(*NotInBetweenCondition); ok {
			return nil
		}
	case Selected:
		if _, ok := conditionPayload.(*SelectedCondition); ok {
			return nil
		}
	case NotSelected:
		if _, ok := conditionPayload.(*NotSelectedCondition); ok {
			return nil
		}
	case Before:
		if _, ok := conditionPayload.(*BeforeCondition); ok {
			return nil
		}
	case After:
		if _, ok := conditionPayload.(*AfterCondition); ok {
			return nil
		}
	}
	return errors.New(fmt.Sprintf("invalid condition for field type %s, condition %v", fieldType, conditionPayload))
}
