package sheet

import "github.com/gin-gonic/gin"

type Controller interface {
	// CreateSheet to create a sheet
	CreateSheet(ctx *gin.Context)
	// GetSheet to get sheet with sheet id
	GetSheet(ctx *gin.Context)
}
