package sheet

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Service interface {
	Create(ctx *gin.Context, tx *gorm.DB, req *CreateSheetRequestDto) (*Sheet, error)
	Get(ctx *gin.Context, tx *gorm.DB, req *GetSheetRequestDto) (*Sheet, error)
}
