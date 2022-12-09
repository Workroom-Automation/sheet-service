package sheet

import (
	"encoding/json"
	"fmt"
	"github.com/leapsquare/sheet-service/pkg/utils"
	"github.com/pkg/errors"
	"time"
)

type LessThanCondition struct {
	Value float64 `json:"value" validate:"required"`
}

type GreaterThanCondition struct {
	Value float64 `json:"value" validate:"required"`
}

type LessThanAndEqualToCondition struct {
	Value float64 `json:"value" validate:"required"`
}

type GreaterThanAndEqualToCondition struct {
	Value float64 `json:"value" validate:"required"`
}

type EqualToCondition struct {
	Value float64 `json:"value" validate:"required"`
}

type NotEqualToCondition struct {
	Value float64 `json:"value" validate:"required"`
}

type BetweenCondition struct {
	UpperLimit float64 `json:"upper_limit" validate:"required"`
	LowerLimit float64 `json:"lower_limit" validate:"required"`
}

type NotInBetweenCondition struct {
	UpperLimit float64 `json:"upper_limit" validate:"required"`
	LowerLimit float64 `json:"lower_limit" validate:"required"`
}

type SelectedCondition struct {
	// TODO decide would store the options in terms of the alphabet or numbers...
	Options []string `json:"options" validate:"required,min=1"`
}

type NotSelectedCondition struct {
	Options []string `json:"options" validate:"required,min=1"`
}

type BeforeCondition struct {
	DateTime time.Time `json:"date_time" validate:"required"`
}

type AfterCondition struct {
	DateTime time.Time `json:"date_time" validate:"required"`
}

type EqualToConditionDate struct {
	DateTime time.Time `json:"date_time" validate:"required"`
}

type BetweenConditionDate struct {
	DateTimeUpperLimit time.Time `json:"upper_limit" validate:"required"`
	DateTimeLowerLimit time.Time `json:"lower_limit" validate:"required"`
}

var ValidTextFieldTypeCondition = []ConditionType{}

var ValidNumberFieldTypeConditions = []ConditionType{
	LessThan,
	GreaterThan,
	LessThanAndEqualTo,
	GreaterThanAndEqualTo,
	EqualTo,
	NotEqualTo,
	Between,
	NotInBetween,
}

var ValidSelectionFieldTypeConditions = []ConditionType{
	Selected,
	NotSelected,
}

var ValidMultiSelectionFieldTypeConditions = []ConditionType{
	Selected,
	NotSelected,
}

var ValidDateSelectionFieldTypeConditions = []ConditionType{
	Before,
	After,
	Between,
	EqualTo,
}

func IsValidConditionForFieldType(fieldType FieldType, conditionType ConditionType) bool {
	switch fieldType {
	case Number:
		for _, validConditionType := range ValidNumberFieldTypeConditions {
			if validConditionType == conditionType {
				return true
			}
		}
	case Selection:
		for _, validConditionType := range ValidSelectionFieldTypeConditions {
			if validConditionType == conditionType {
				return true
			}
		}
	case MultiSelection:
		for _, validConditionType := range ValidMultiSelectionFieldTypeConditions {
			if validConditionType == conditionType {
				return true
			}
		}
	case DateSelection:
		for _, validConditionType := range ValidDateSelectionFieldTypeConditions {
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
	jsonString, _ := json.Marshal(conditionPayload)
	validate := utils.GetJsonValidator()
	switch conditionType {
	case LessThan:
		var lessThan LessThanCondition
		err := json.Unmarshal(jsonString, &lessThan)
		if err != nil {
			break
		}
		err = validate.Struct(lessThan)
		if err == nil {
			return nil
		}
	case GreaterThan:
		var greaterThan GreaterThanCondition
		err := json.Unmarshal(jsonString, &greaterThan)
		if err != nil {
			break
		}
		err = validate.Struct(greaterThan)
		if err == nil {
			return nil
		}
	case LessThanAndEqualTo:
		var lessThanAndEqualTo LessThanAndEqualToCondition
		err := json.Unmarshal(jsonString, &lessThanAndEqualTo)
		if err != nil {
			break
		}
		err = validate.Struct(lessThanAndEqualTo)
		if err == nil {
			return nil
		}
	case GreaterThanAndEqualTo:
		var greaterThanAndEqualTo GreaterThanAndEqualToCondition
		err := json.Unmarshal(jsonString, &greaterThanAndEqualTo)
		if err != nil {
			break
		}
		err = validate.Struct(greaterThanAndEqualTo)
		if err == nil {
			return nil
		}
	case EqualTo:
		if fieldType == DateSelection {
			var equalTo EqualToConditionDate
			err := json.Unmarshal(jsonString, &equalTo)
			if err != nil {
				break
			}
			err = validate.Struct(equalTo)
			if err == nil {
				return nil
			}
		} else {
			var equalTo EqualToCondition
			err := json.Unmarshal(jsonString, &equalTo)
			if err != nil {
				break
			}
			err = validate.Struct(equalTo)
			if err == nil {
				return nil
			}
		}
	case NotEqualTo:
		var notEqualTo NotEqualToCondition
		err := json.Unmarshal(jsonString, &notEqualTo)
		if err != nil {
			break
		}
		err = validate.Struct(notEqualTo)
		if err == nil {
			return nil
		}
	case Between:
		if fieldType == DateSelection {
			var between BetweenConditionDate
			err := json.Unmarshal(jsonString, &between)
			if err != nil {
				break
			}
			err = validate.Struct(between)
			if err == nil {
				return nil
			}
		} else {
			var between BetweenCondition
			err := json.Unmarshal(jsonString, &between)
			if err != nil {
				break
			}
			err = validate.Struct(between)
			if err == nil {
				return nil
			}
		}
	case NotInBetween:
		var notInBetween NotInBetweenCondition
		err := json.Unmarshal(jsonString, &notInBetween)
		if err != nil {
			break
		}
		err = validate.Struct(notInBetween)
		if err == nil {
			return nil
		}
	case Selected:
		var selected SelectedCondition
		err := json.Unmarshal(jsonString, &selected)
		if err != nil {
			break
		}
		err = validate.Struct(selected)
		if err == nil {
			return nil
		}
	case NotSelected:
		var notSelected NotSelectedCondition
		err := json.Unmarshal(jsonString, &notSelected)
		if err != nil {
			break
		}
		err = validate.Struct(notSelected)
		if err == nil {
			return nil
		}
	case Before:
		var before BeforeCondition
		err := json.Unmarshal(jsonString, &before)
		if err != nil {
			break
		}
		err = validate.Struct(before)
		if err == nil {
			return nil
		}
	case After:
		var after AfterCondition
		err := json.Unmarshal(jsonString, &after)
		if err != nil {
			break
		}
		err = validate.Struct(after)
		if err == nil {
			return nil
		}
	}
	return errors.New(fmt.Sprintf("invalid condition for field type %s, condition %v", fieldType, conditionPayload))
}
