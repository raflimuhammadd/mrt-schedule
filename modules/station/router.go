package station

import (
	"net/http"
	"strings"

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

	station.GET("/id/:id", func(c *gin.Context) {
		CheckSchedulesByStation(c, stationService) 
	})

	station.GET("/:slug/schedules", func(c *gin.Context) {
		GetStationSchedules(c, stationService)
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

func GetStationSchedules(c *gin.Context, service Service) {
	slug := c.Param("slug")

	if slug == "" {
		c.JSON(
			http.StatusBadRequest,
			response.APIResponse{
				Success: false,
				Message: "Station slug is required",
				Data: nil,
			},
		)
		return
	}

	data, err := service.GetStationScheduleBySlug(slug)
	if err != nil {
		errorMsg := err.Error()
		statusCode := http.StatusInternalServerError

		if strings.Contains(errorMsg, "not found") {
			statusCode = http.StatusNotFound
		} else if strings.Contains(errorMsg, "required") {
			statusCode = http.StatusBadRequest
		}

		c.JSON(
			statusCode,
			response.APIResponse{
				Success: false,
				Message: errorMsg,
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
			Data: data,
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