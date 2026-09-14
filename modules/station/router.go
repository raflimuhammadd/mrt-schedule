package station

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/raflimuhammadd/mrt-schedule/common/response"
)

func Initiate(router *gin.RouterGroup) {
	stationService := NewService()

	station := router.Group("/stations")
	station.GET("", func(c *gin.Context) {
		// services
		GetAllStation(c, stationService)
	})

	station.GET("/:id", func(c *gin.Context) {
		CheckSchedulesByStation(c, stationService) 
	})
}

func GetAllStation(c *gin.Context, service Service) {
	datas, err := service.GetAllStation()
	if err != nil {
		// handle error
		c.JSON(
			http.StatusBadRequest,
			response.APIResponse{
				Success: false,
				Message: err.Error(),
				Data: nil,
			},
		)
		return
	}

	// response
	c.JSON(
		http.StatusOK,
		response.APIResponse{
			Success: true,
			Message: "successfully get all stations",
			Data: datas,
		},
	)
} 

func CheckSchedulesByStation(c *gin.Context, service Service) {
	id := c.Param("id")

	datas, err := service.CheckScheduleByStation(id)
	if err != nil {
		// handle error
		c.JSON(
			http.StatusBadRequest,
			response.APIResponse{
				Success: false,
				Message: err.Error(),
				Data: nil,
			},
		)
		return
	}

	c.JSON(
		http.StatusOK,
		response.APIResponse{
			Success: true,
			Message: "successfully get schedules",
			Data: datas,
		},
	)
}