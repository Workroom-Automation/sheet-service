package sheet

import "github.com/gin-gonic/gin"

type Service interface {
	Create(ctx *gin.Context)
}
