package sheet

import (
	"github.com/leapsquare/sheet-service/pkg/logger"
	"gorm.io/gorm"
)

type repo struct {
	logger logger.Logger
}

func NewRepository(logger logger.Logger) Repository {
	return &repo{
		logger: logger,
	}
}

func (r *repo) Create(tx *gorm.DB, sheet *Sheet) error {
	return tx.Model(&Sheet{}).Create(&sheet).Error
}

func (r *repo) Get(tx *gorm.DB, sheetId int64) (*Sheet, error) {
	var sheet *Sheet
	if err := tx.Model(&Sheet{}).First(&sheet, sheetId).Error; err != nil {
		return nil, err
	}
	return sheet, nil
}
