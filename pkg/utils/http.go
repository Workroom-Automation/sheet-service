package utils

import (
	"github.com/gin-gonic/gin"
)

const (
	ctxUserId = "X-User-Id"
)

func GetConfigPath(configPath string) string {
	return "./config/default"
}

func GetRequestCtx(in *gin.Context) *gin.Context {
	//data := in.Values().Get(auth.CtxValidatedData).(*auth.ValidatedData)
	//ctx = context.WithValue(ctx, ctxUserId, data.User.ID)
	return in
}

func GetUserId(ctx *gin.Context) string {
	return ctx.Request.Header.Get(ctxUserId)
}
