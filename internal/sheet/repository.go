package sheet

import (
	"gorm.io/gorm"
)

type Repository interface {
	Create(tx *gorm.DB, sheet *Sheet) error
	Get(tx *gorm.DB, sheetId int64) (*Sheet, error)
	Update(tx *gorm.DB, sheet *Sheet) error
}
