package sheet

import (
	"encoding/json"
	"fmt"
	"github.com/leapsquare/sheet-service/pkg/utils"
	"github.com/pkg/errors"
)

type SendMailAction struct {
	To      []*string `json:"to" validate:"required,min=1"`
	Body    string    `json:"body" validate:"required"`
	Subject string    `json:"subject" validate:"required"`
}

func ValidateTriggerActionForActionType(actionType ActionType, action interface{}) error {
	jsonString, _ := json.Marshal(action)
	validate := utils.GetJsonValidator()
	switch actionType {
	case SendMail:
		var sendMailAction SendMailAction
		err := json.Unmarshal(jsonString, &sendMailAction)
		if err != nil {
			break
		}
		err = validate.Struct(sendMailAction)
		if err == nil {
			return nil
		}
	}
	return errors.New(fmt.Sprintf("invalid action payload for action type %s, action payload %v", actionType, action))
}
