package sheet

import (
	"github.com/gin-gonic/gin"
	"github.com/leapsquare/sheet-service/internal/middleware"
)

func MapRoutesV1(party *gin.RouterGroup, handlers Controller, m *middleware.Middleware) {
	party.Use(m.EnrichContextFromAuthHeader)

	appParty := party.Group("/sheet")
	appParty.GET("/", handlers.GetSheet)
	appParty.POST("/", handlers.CreateSheet)
}
