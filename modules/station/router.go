package station

import (
	"mrt-schedules-go/common/response"
	"net/http"
	"github.com/gin-gonic/gin"
)

func Initiate(router *gin.RouterGroup){
	stationService := NewService()
	// stationController := controllers.StationController
	station := router.Group("/station")
	station.GET("", func(ctx *gin.Context) {
		GetAllStation(ctx, stationService)
	})

	station.GET("/:id", func(ctx *gin.Context) {
		CheckSchedule (ctx, stationService)
	})

	station.GET("/retail/:id", func(ctx *gin.Context) {
		CheckRetail(ctx, stationService)
	})

	station.GET("/facility/:id", func(ctx *gin.Context) {
		CheckFacility(ctx, stationService)
	})
}

func GetAllStation(ctx *gin.Context, service Service){
	datas, err := service.GetAllStation()

	if err != nil {
		ctx.JSON(http.StatusBadRequest, response.APIResponse{
			Success: false,
			Message: err.Error(),
			Data: nil,
		})
		return
	}

	ctx.JSON(http.StatusOK, response.APIResponse{
		Success: true,
		Message: "success",
		Data: datas,
	})
}

func CheckSchedule(ctx *gin.Context, service Service){
	id := ctx.Param("id")
	datas, err := service.CheckSchedule(id)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, response.APIResponse{
			Success: false,
			Message: err.Error(),
			Data: nil,
		})
		return
	}

	ctx.JSON(http.StatusOK, response.APIResponse{
		Success: true,
		Message: "success",
		Data: datas,
	})
}

func CheckRetail(ctx *gin.Context, service Service){
	id := ctx.Param("id")
	datas, err := service.CheckRetail(id)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, response.APIResponse{
			Success: false,
			Message: err.Error(),
			Data: nil,
		})
		return
	}

	ctx.JSON(http.StatusOK, response.APIResponse{
		Success: true,
		Message: "success",
		Data: datas,
	})
}

func CheckFacility(ctx *gin.Context, service Service){
	id := ctx.Param("id")
	datas, err := service.CheckFacility(id)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, response.APIResponse{
			Success: false,
			Message: err.Error(),
			Data: nil,
		})
		return
	}

	ctx.JSON(http.StatusOK, response.APIResponse{
		Success: true,
		Message: "success",
		Data: datas,
	})
}