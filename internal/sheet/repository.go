package sheet

import (
	rqp "github.com/timsolov/rest-query-parser"
	"gorm.io/gorm"
)

type Repository interface {
	Create(tx *gorm.DB, sheet *Sheet) error
	Get(tx *gorm.DB, sheetId int64) (*Sheet, error)
	Update(tx *gorm.DB, sheet *Sheet) error
	List(tx *gorm.DB, q *rqp.Query) ([]*TrimmedSheet, error)
}
