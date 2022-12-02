package sheet

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"github.com/leapsquare/sheet-service/internal/model"
	"github.com/pkg/errors"
)

type Properties struct {
	Type               string      `json:"type"`
	Title              string      `json:"title"`
	Description        string      `json:"description"`
	IsRequired         *bool       `json:"is_required"`
	RequireObservation *bool       `json:"require_observation"`
	HasAttachment      *bool       `json:"has_attachment"`
	Data               interface{} `json:"data"`
}

type Trigger struct {
	Id            int64       `json:"id" gorm:"column:id"`
	Name          string      `json:"name" gorm:"column:name"`
	ConditionType string      `json:"condition_type" gorm:"column:condition_type"`
	Condition     interface{} `json:"condition" gorm:"column:condition"`
	ActionType    string      `json:"action_type" gorm:"column:action_type"`
	Action        interface{} `json:"action" gorm:"column:action"`
}

type Field struct {
	Id         int64       `json:"id" gorm:"column:id"`
	Properties *Properties `json:"properties" gorm:"column:properties"` // it can be a json
	Triggers   []*Trigger  `json:"triggers" gorm:"column:triggers"`
}

type Section struct {
	Id          int64    `json:"id" gorm:"column:id"`
	Name        string   `json:"name" gorm:"column:name"`
	Description string   `json:"description" gorm:"column:description"`
	Fields      []*Field `json:"fields" gorm:"column:fields"`
}

type SheetSections struct {
	Sections []*Section `json:"sections"`
}

type Sheet struct {
	*model.BaseModel
	Id            int64          `json:"id" gorm:"column:id"`
	Name          string         `json:"name" gorm:"column:name"`
	Description   string         `json:"description" gorm:"column:description"`
	ExternalCode  string         `json:"external_code" gorm:"column:external_code"` // TODO verify this must be unique
	ApplicationId string         `json:"application_id" gorm:"column:application_id"`
	AssetId       string         `json:"asset_id" gorm:"column:asset_id"`
	ProcessId     string         `json:"process_id" gorm:"column:process_id"`
	IsActive      *bool          `json:"is_active" gorm:"column:is_active"`
	SheetSections *SheetSections `json:"sheet_sections" gorm:"column:sheet_sections"`
}

// Scan implements the sql.Scanner interface
func (s *SheetSections) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New(fmt.Sprint("Failed to unmarshal JSONB value:", value))
	}
	err := json.Unmarshal(bytes, s)
	return err
}

func (s *SheetSections) Value() (driver.Value, error) {
	message, err := json.Marshal(s)
	return message, err
}

type TrimmedSheet struct {
	*model.BaseModel
	Id            int64  `json:"id" gorm:"column:id"`
	Name          string `json:"name" gorm:"column:name"`
	Description   string `json:"description" gorm:"column:description"`
	ExternalCode  string `json:"external_code" gorm:"column:external_code"` // TODO verify this must be unique
	ApplicationId string `json:"application_id" gorm:"column:application_id"`
	AssetId       string `json:"asset_id" gorm:"column:asset_id"`
	ProcessId     string `json:"process_id" gorm:"column:process_id"`
	IsActive      *bool  `json:"is_active" gorm:"column:is_active"`
}
