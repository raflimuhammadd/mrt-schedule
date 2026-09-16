package journey

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/raflimuhammadd/mrt-schedule/common/response"
)

func Initiate(router *gin.RouterGroup) {
	journeyService := NewService()

	journey := router.Group("/journeys")
	
	// Main journey planning endpoint
	journey.GET("/plan", func(c *gin.Context) {
		PlanJourney(c, journeyService)
	})

	// Single fare calculation endpoint
	journey.GET("/fare", func(c *gin.Context) {
		GetFare(c, journeyService)
	})

	// Single duration calculation endpoint
	journey.GET("/duration", func(c *gin.Context) {
		GetDuration(c, journeyService)
	})
}

// PlanJourney handle journey planning request
func PlanJourney(c *gin.Context, service Service) {
	// Get query parameters
	fromSlug := c.Query("from")
	toSlug := c.Query("to")

	// Validate parameters
	if fromSlug == "" {
		c.JSON(
			http.StatusBadRequest,
			response.APIResponse{
				Success: false,
				Message: "Parameter 'from' is required",
				Data:    nil,
			},
		)
		return
	}

	if toSlug == "" {
		c.JSON(
			http.StatusBadRequest,
			response.APIResponse{
				Success: false,
				Message: "Parameter 'to' is required",
				Data:    nil,
			},
		)
		return
	}

	// Validate: same station?
	if strings.EqualFold(fromSlug, toSlug) {
		c.JSON(
			http.StatusBadRequest,
			response.APIResponse{
				Success: false,
				Message: "From and to stations cannot be the same",
				Data:    nil,
			},
		)
		return
	}

	// Call service
	data, err := service.PlanJourney(fromSlug, toSlug)
	if err != nil {
		errorMsg := err.Error()
		statusCode := http.StatusInternalServerError

		if strings.Contains(errorMsg, "not found") {
			statusCode = http.StatusNotFound
		} else if strings.Contains(errorMsg, "cannot be the same") {
			statusCode = http.StatusBadRequest
		}

		c.JSON(
			statusCode,
			response.APIResponse{
				Success: false,
				Message: errorMsg,
				Data:    nil,
			},
		)
		return
	}

	// Success response
	c.JSON(
		http.StatusOK,
		response.APIResponse{
			Success: true,
			Message: "successfully plan journey",
			Data:    data,
		},
	)
}

// GetFare handle fare calculation request
func GetFare(c *gin.Context, service Service) {
	fromSlug := c.Query("from")
	toSlug := c.Query("to")

	if fromSlug == "" {
		c.JSON(http.StatusBadRequest, response.APIResponse{
			Success: false,
			Message: "Parameter 'from' is required",
			Data:    nil,
		})
		return
	}

	if toSlug == "" {
		c.JSON(http.StatusBadRequest, response.APIResponse{
			Success: false,
			Message: "Parameter 'to' is required",
			Data:    nil,
		})
		return
	}

	data, err := service.CalculateFare(fromSlug, toSlug)
	if err != nil {
		errorMsg := err.Error()
		statusCode := http.StatusInternalServerError

		if strings.Contains(errorMsg, "not found") {
			statusCode = http.StatusNotFound
		}

		c.JSON(statusCode, response.APIResponse{
			Success: false,
			Message: errorMsg,
			Data:    nil,
		})
		return
	}

	c.JSON(http.StatusOK, response.APIResponse{
		Success: true,
		Message: "successfully calculate fare",
		Data:    data,
	})
}

// GetDuration handle duration calculation request
func GetDuration(c *gin.Context, service Service) {
	fromSlug := c.Query("from")
	toSlug := c.Query("to")

	if fromSlug == "" {
		c.JSON(http.StatusBadRequest, response.APIResponse{
			Success: false,
			Message: "Parameter 'from' is required",
			Data:    nil,
		})
		return
	}

	if toSlug == "" {
		c.JSON(http.StatusBadRequest, response.APIResponse{
			Success: false,
			Message: "Parameter 'to' is required",
			Data:    nil,
		})
		return
	}

	data, err := service.CalculateDuration(fromSlug, toSlug)
	if err != nil {
		errorMsg := err.Error()
		statusCode := http.StatusInternalServerError

		if strings.Contains(errorMsg, "not found") {
			statusCode = http.StatusNotFound
		}

		c.JSON(statusCode, response.APIResponse{
			Success: false,
			Message: errorMsg,
			Data:    nil,
		})
		return
	}

	c.JSON(http.StatusOK, response.APIResponse{
		Success: true,
		Message: "successfully calculate duration",
		Data:    data,
	})
}