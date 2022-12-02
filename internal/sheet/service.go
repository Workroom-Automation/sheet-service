package sheet

import (
	"github.com/gin-gonic/gin"
	rqp "github.com/timsolov/rest-query-parser"
	"gorm.io/gorm"
)

type Service interface {
	Create(ctx *gin.Context, tx *gorm.DB, req *CreateSheetRequestDto) (*Sheet, error)
	Get(ctx *gin.Context, tx *gorm.DB, sheetId int64) (*Sheet, error)
	GetSheetAuthoringPlatformResources(ctx *gin.Context, tx *gorm.DB) (*SheetAuthoringPlatformResources, error)
	Update(ctx *gin.Context, tx *gorm.DB, req *UpdateSheetRequestDto) (*Sheet, error)
	List(ctx *gin.Context, tx *gorm.DB, q *rqp.Query) ([]*TrimmedSheet, error)
}
