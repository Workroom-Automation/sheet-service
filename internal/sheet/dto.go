package sheet

type FieldType string

var (
	Text           FieldType = "TEXT"
	Number         FieldType = "NUMBER"
	Selection      FieldType = "SELECTION"
	MultiSelection FieldType = "MULTI_SELECTION"
	DateSelection  FieldType = "DATE_SELECTION"
	//RangePicker   FieldType = ""
	//UserSelect    FieldType = ""
	//MediaUpload   FieldType = ""
)

type ActionType string

var (
	SendMail ActionType = "SEND_MAIL"
)

type ConditionType string

var (
	LessThan              ConditionType = "LT"
	GreaterThan           ConditionType = "GT"
	LessThanAndEqualTo    ConditionType = "LTE"
	GreaterThanAndEqualTo ConditionType = "GTE"
	EqualTo               ConditionType = "EQ"
	NotEqualTo            ConditionType = "NEQ"
	Between               ConditionType = "BW"
	NotInBetween          ConditionType = "NBW"
	Selected              ConditionType = "SE"
	NotSelected           ConditionType = "NSE"
	Before                ConditionType = "BF"
	After                 ConditionType = "AF"
)

type CreateSheetRequestDto struct {
	Name          string        `json:"name" binding:"required"`
	Description   string        `json:"description" bidding:"required"`
	ExternalCode  string        `json:"external_code" bidding:"required"`
	ApplicationId string        `json:"application_id" bidding:"required"`
	AssetId       string        `json:"asset_id" bidding:"required"`
	ProcessId     string        `json:"process_id" bidding:"required"`
	IsActive      *bool         `json:"is_active" bidding:"required"`
	Sections      []*SectionDto `json:"sections" binding:"dive"`
}

type PropertiesDto struct {
	Type          FieldType   `json:"type" binding:"required"`
	Title         string      `json:"title" binding:"required"`
	Description   string      `json:"description" binding:"required"`
	IsRequired    *bool       `json:"is_required"`
	HasAttachment *bool       `json:"has_attachment"`
	FieldFormData interface{} `json:"field_form_data"`
}

type TriggerDto struct {
	Name          string        `json:"name" binding:"required"`
	ConditionType ConditionType `json:"condition_type"`
	Condition     interface{}   `json:"condition" binding:"required"`
	ActionType    ActionType    `json:"action_type"`
	Action        interface{}   `json:"action" binding:"required"`
}

type FieldDto struct {
	Properties *PropertiesDto `json:"properties" binding:"required,dive"`
	Triggers   []*TriggerDto  `json:"triggers" binding:"dive"`
}

type SectionDto struct {
	Name        string      `json:"name" binding:"required"`
	Description string      `json:"description" binding:"required"`
	Fields      []*FieldDto `json:"fields"`
}

func (t *TriggerDto) ToTrigger() *Trigger {
	return &Trigger{
		Name:          t.Name,
		Condition:     t.Condition,
		Action:        t.Action,
		ConditionType: string(t.ConditionType),
		ActionType:    string(t.ActionType),
	}
}
func (p *PropertiesDto) ToProperties() *Properties {
	return &Properties{
		Type:          string(p.Type),
		Title:         p.Title,
		Description:   p.Description,
		IsRequired:    p.IsRequired,
		HasAttachment: p.HasAttachment,
		Data:          p.FieldFormData,
	}
}

func (f *FieldDto) ToField() *Field {
	properties := f.Properties.ToProperties()
	var triggers []*Trigger
	for _, ele := range f.Triggers {
		triggers = append(triggers, ele.ToTrigger())
	}
	return &Field{
		Properties: properties,
		Triggers:   triggers,
	}
}

func (s *SectionDto) ToSection() *Section {
	var fields []*Field
	for _, ele := range s.Fields {
		fields = append(fields, ele.ToField())
	}
	return &Section{
		Name:        s.Name,
		Description: s.Description,
		Fields:      fields,
	}
}

func (c *CreateSheetRequestDto) ToSheet() *Sheet {
	isActive := true
	var sheetSections = &SheetSections{}
	var sections []*Section
	for _, ele := range c.Sections {
		sections = append(sections, ele.ToSection())
	}
	sheetSections.Sections = sections
	return &Sheet{
		Name:          c.Name,
		Description:   c.Description,
		ExternalCode:  c.ExternalCode,
		ApplicationId: c.ApplicationId,
		AssetId:       c.AssetId,
		ProcessId:     c.ProcessId,
		IsActive:      &isActive,
		SheetSections: sheetSections,
	}
}

type FieldResource struct {
	TriggerConditions []ConditionType `json:"trigger_conditions"`
}

type SheetAuthoringPlatformResources struct {
	FieldResources map[FieldType]FieldResource `json:"field_resources"`
}
