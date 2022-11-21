package sheet

import (
	"github.com/gin-gonic/gin"
	"github.com/leapsquare/sheet-service/config"
	"github.com/leapsquare/sheet-service/pkg/logger"
	"github.com/leapsquare/sheet-service/pkg/utils"
	"net/http"
)

type controller struct {
	svc    Service
	logger logger.Logger
	cfg    *config.Config
}

func NewController(cfg *config.Config, logger logger.Logger, svc Service) Controller {
	return &controller{
		cfg:    cfg,
		logger: logger,
		svc:    svc,
	}
}

func (c controller) CreateSheet(ctx *gin.Context) {
	request := CreateSheetRequestDto{}
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.String(http.StatusBadRequest, err.Error())
		return
	}
	customCtx := utils.GetRequestCtx(ctx)
	sheet, err := c.svc.Create(customCtx, nil, &request)
	if err != nil {
		ctx.String(http.StatusBadRequest, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, sheet)
}

func (c controller) GetSheet(ctx *gin.Context) {
	val, err := utils.GetQueryInt64(ctx, "id")
	if err != nil {
		ctx.String(http.StatusBadRequest, err.Error())
		return
	}
	customCtx := utils.GetRequestCtx(ctx)
	sheet, err := c.svc.Get(customCtx, nil, val)
	if err != nil {
		ctx.String(http.StatusBadRequest, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, sheet)
}

func (c controller) GetSheetAuthoringPlatformResources(ctx *gin.Context) {
	customCtx := utils.GetRequestCtx(ctx)
	resources, err := c.svc.GetSheetAuthoringPlatformResources(customCtx, nil)
	if err != nil {
		ctx.String(http.StatusBadRequest, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, resources)
}

func (c controller) UpdateSheet(ctx *gin.Context) {
	request := UpdateSheetRequestDto{}
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.String(http.StatusBadRequest, err.Error())
		return
	}
	customCtx := utils.GetRequestCtx(ctx)
	sheet, err := c.svc.Update(customCtx, nil, &request)
	if err != nil {
		ctx.String(http.StatusBadRequest, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, sheet)
}

func (c controller) LogSheet(ctx *gin.Context) {

}
