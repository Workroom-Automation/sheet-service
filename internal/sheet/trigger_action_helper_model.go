package sheet

import (
	"fmt"
	"github.com/pkg/errors"
)

type SendMailAction struct {
	To      []string `json:"to"`
	Body    string   `json:"body"`
	Subject string   `json:"subject"`
}

func ValidateTriggerActionForActionType(actionType ActionType, action interface{}) error {
	switch actionType {
	case SendMail:
		if _, ok := action.(*SendMailAction); ok {
			return nil
		}
	}
	return errors.New(fmt.Sprintf("invalid action payload for action type %s, action payload %v", actionType, action))
}
