package sheet

import (
	"github.com/gin-gonic/gin"
	"github.com/leapsquare/sheet-service/config"
	"github.com/leapsquare/sheet-service/pkg/logger"
	"gorm.io/gorm"
)

/**
Validations:
1. a valid mapping between field type and actions and conditions.....
2. "Data" validation for the field type...
*/

type service struct {
	repo   Repository
	logger logger.Logger
	cfg    *config.Config
	db     *gorm.DB
}

func NewService(logger logger.Logger, db *gorm.DB, cfg *config.Config, repo Repository) Service {
	return &service{
		logger: logger,
		db:     db,
		cfg:    cfg,
		repo:   repo,
	}
}

func (s *service) DbWithContext(ctx *gin.Context, in *gorm.DB) *gorm.DB {
	if in == nil {
		return s.db.WithContext(ctx)
	}
	return in
}

func (s *service) Create(ctx *gin.Context, tx *gorm.DB, req *CreateSheetRequestDto) (*Sheet, error) {
	/**
	Validate the app, asset , process via the api calls......
	*/
	tx = s.DbWithContext(ctx, tx)
	sections := req.Sections
	// Start the section validations...
	for _, section := range sections {
		fields := section.Fields
		// start the field validations...
		for _, field := range fields {
			// field type will be the main component on which trigger and condition will be decided
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
	sheet := req.ToSheet()
	if err := s.repo.Create(tx, sheet); err != nil {
		return nil, err
	}
	return sheet, nil
}

func (s *service) Get(ctx *gin.Context, tx *gorm.DB, sheetId int64) (*Sheet, error) {
	tx = s.DbWithContext(ctx, tx)
	return s.repo.Get(tx, sheetId)
}
