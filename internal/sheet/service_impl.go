package sheet

import (
	"github.com/gin-gonic/gin"
)

/**
Validations:
1. a valid mapping between field type and actions and conditions.....
2. "Data" validation for the field type...

*/

// CreateSheet something
func Create(ctx *gin.Context, req *CreateSheetRequestDto) (*Sheet, error) {

	// TODO::
	/**
	Validate the app, asset , process via the api calls......
	*/
	sections := req.Sections
	// Start the section validations...
	for _, section := range sections {
		fields := section.Fields
		// start the field validations...
		for _, field := range fields {
			fieldType := field.Properties.Type
			if err := ValidateFieldFormDataForFieldType(fieldType, field.Properties.FieldFormData); err != nil {
				return nil, err
			}
			// validating the triggers....
			for _, trigger := range field.Triggers {
				actionType := trigger.ActionType
				if err := ValidateTriggerActionForActionType(actionType, trigger.Action); err != nil {
					return nil, err
				}
				if err := ValidateTriggerConditionForFieldType(fieldType, trigger.ConditionType, trigger.Condition); err != nil {
					return nil, err
				}
			}
		}
	}
	return req.ToSheet(), nil
}

func Get(ctx *gin.Context, req *GetSheetRequestDto) (*Sheet, error) {
	return &Sheet{}, nil
}
