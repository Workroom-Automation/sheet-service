package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/leapsquare/sheet-service/config"
	"github.com/leapsquare/sheet-service/pkg/logger"
)

type Middleware struct {
	cfg    *config.Config
	logger logger.Logger
}

func NewMiddleware(cfg *config.Config, logger logger.Logger) *Middleware {
	return &Middleware{cfg: cfg, logger: logger}
}

func (m Middleware) EnrichContextFromAuthHeader(ctx *gin.Context) {
	//data := ctx.Values().Get("validated-data").(*auth.ValidatedData)
	//if data.User.ID == "" {
	//	return
	//}
	ctx.Next()
}
