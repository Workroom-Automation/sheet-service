package sheet

import "github.com/leapsquare/sheet-service/internal/model"

type Properties struct {
	Type          string      `json:"type"`
	Title         string      `json:"title"`
	Description   string      `json:"description"`
	IsRequired    *bool       `json:"is_required"`
	HasAttachment *bool       `json:"has_attachment"`
	Data          interface{} `json:"data"`
}

type Trigger struct {
	Id            int64       `json:"id" gorm:"column:id"`
	FieldId       int64       `json:"field_id" gorm:"column:field_id"`
	Name          string      `json:"name" gorm:"column:name"`
	ConditionType string      `json:"condition_type" gorm:"column:condition_type"`
	Condition     interface{} `json:"condition" gorm:"column:condition"`
	ActionType    string      `json:"action_type" gorm:"column:action_type"`
	Action        interface{} `json:"action" gorm:"column:action"`
}

type Field struct {
	Id         int64       `json:"id" gorm:"column:id"`
	SectionId  int64       `json:"section_id" gorm:"column:section_id"`
	Properties *Properties `json:"properties" gorm:"column:properties"` // it can be a json
	Triggers   []*Trigger  `json:"triggers" gorm:"column:triggers"`
}

type Section struct {
	Id          int64    `json:"id" gorm:"column:id"`
	SheetId     int64    `json:"sheet_id" gorm:"column:sheet_id"`
	Name        string   `json:"name" gorm:"column:name"`
	Description string   `json:"description" gorm:"column:description"`
	Fields      []*Field `json:"fields" gorm:"column:fields"`
}

// https://stackoverflow.com/questions/58633251/unmarshal-json-array-of-object-obtained-from-postgresql

type Sheet struct {
	*model.BaseModel
	Id            int64      `json:"id" gorm:"column:id"`
	Name          string     `json:"name" gorm:"column:name"`
	Description   string     `json:"description" gorm:"column:description"`
	ExternalCode  string     `json:"external_code" gorm:"column:external_code"` // TODO verify this must be unique
	ApplicationId string     `json:"application_id" gorm:"column:application_id"`
	AssetId       string     `json:"asset_id" gorm:"column:asset_id"`
	ProcessId     string     `json:"process_id" gorm:"column:process_id"`
	IsActive      *bool      `json:"is_active" gorm:"column:is_active"`
	Sections      []*Section `json:"sections" gorm:"column:sections"`
}
