package sheet

import (
	"github.com/leapsquare/sheet-service/pkg/logger"
	"github.com/pkg/errors"
	rqp "github.com/timsolov/rest-query-parser"
	"gorm.io/gorm"
)

const (
	SheetTable = "sheets"
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
	return tx.Table(SheetTable).Create(&sheet).Error
}

func (r *repo) Get(tx *gorm.DB, sheetId int64) (*Sheet, error) {
	var sheet *Sheet
	if err := tx.Table(SheetTable).First(&sheet, sheetId).Error; err != nil {
		return nil, err
	}
	return sheet, nil
}

func (r *repo) Update(tx *gorm.DB, sheet *Sheet) error {
	result := tx.Table(SheetTable).Where("id = ?", sheet.Id).Updates(&sheet)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("update failed")
	}
	return nil
}

func (r *repo) List(tx *gorm.DB, q *rqp.Query) ([]*TrimmedSheet, error) {
	var sheets = []*TrimmedSheet{}
	err := tx.Table(SheetTable).Where(q.Where(), q.Args()...).Find(&sheets).Limit(q.Limit).Offset(q.Offset).Error
	return sheets, err
}
