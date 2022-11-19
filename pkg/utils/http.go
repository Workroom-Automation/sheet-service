package utils

import (
	"github.com/gin-gonic/gin"
	"strconv"
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

func GetQueryInt64(ctx *gin.Context, key string) (int64, error) {
	val := ctx.Query(key)
	return strconv.ParseInt(val, 10, 64)
}
