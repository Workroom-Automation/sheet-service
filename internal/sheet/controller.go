package sheet

import "github.com/gin-gonic/gin"

type Controller interface {
	// CreateSheet to create a sheet
	CreateSheet(ctx *gin.Context)
	// GetSheet to get sheet with sheet id
	GetSheet(ctx *gin.Context)
	// GetSheetAuthoringPlatformResources to fetch every information related to sheet authoring module
	GetSheetAuthoringPlatformResources(ctx *gin.Context)
	// UpdateSheet to update the pre authored sheet
	UpdateSheet(ctx *gin.Context)
	// LogSheet to log a sheet and have all the validations and triggers....
	LogSheet(ctx *gin.Context)
	// GetSheets to get the list of the sheet
	GetSheets(ctx *gin.Context)
}
