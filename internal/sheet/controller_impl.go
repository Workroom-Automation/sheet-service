package sheet

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/leapsquare/sheet-service/config"
	"github.com/leapsquare/sheet-service/pkg/logger"
	"github.com/leapsquare/sheet-service/pkg/utils"
	rqp "github.com/timsolov/rest-query-parser"
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
	val, err := utils.GetPathInt64(ctx, "id")
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

func (c controller) GetSheets(ctx *gin.Context) {
	fmt.Println(ctx.Request.URL.Query())
	q, err := rqp.NewParse(ctx.Request.URL.Query(), rqp.Validations{
		"limit:required": rqp.MinMax(1, 50),
		"application_id": nil,
		"asset_id":       nil,
		"process_id":     nil,
	})
	if err != nil {
		ctx.String(http.StatusBadRequest, err.Error())
		return
	}
	// add is active to true
	customCtx := utils.GetRequestCtx(ctx)
	sheets, err := c.svc.List(customCtx, nil, q)
	if err != nil {
		ctx.String(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, sheets)
}
