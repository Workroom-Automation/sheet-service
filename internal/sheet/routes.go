package sheet

import (
	"github.com/gin-gonic/gin"
	"github.com/leapsquare/sheet-service/internal/middleware"
)

func MapRoutesV1(party *gin.RouterGroup, handlers Controller, m *middleware.Middleware) {
	party.Use(m.EnrichContextFromAuthHeader)

	sheetParty := party.Group("/sheet")
	sheetParty.GET("/", handlers.GetSheet)
	sheetParty.POST("/", handlers.CreateSheet)
	sheetParty.PATCH("/", handlers.UpdateSheet)
	sheetParty.GET("/canvas", handlers.GetSheetAuthoringPlatformResources)
}
