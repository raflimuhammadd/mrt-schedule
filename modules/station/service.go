package station

import (
	"encoding/json"
	"net/http"
	"os"
	"time"
	"errors"

	"github.com/raflimuhammadd/mrt-schedule/common/client"
)

type Service interface {
	GetAllStation() (response []StationResponse, err error)
	CheckScheduleByStation(id string) (response []StationResponse, err error)
}

type service struct {
	client*http.Client
	baseURL string
}

func NewService() Service {
	return &service{
		client: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

func (s *service) GetAllStation() (response []StationResponse, err error) {
	// layer service
	baseURL := os.Getenv("API_BASE_URL")
	if baseURL == "" {
		return nil, errors.New("API_BASE_URL is required")
	}

	url := baseURL + "/datum"

	byteResponse, err := client.DoRequest(s.client, url)
	if err != nil {
		return
	}

	var apiResp StationAPIResponse
	err = json.Unmarshal(byteResponse, &apiResp)
	if err != nil {
		return
	}

	for _, item := range apiResp.Data {
		response = append(response, StationResponse{
			Id: item.Id,
			Name: item.Name,
			Description: item.Description,
		})
	}
	
	return 
}

func (s *service) CheckScheduleByStation(id string) (response []StationResponse, err error) {
	url := "https://beweb-dev.jakartamrt.co.id/middleware/api/datum"

	byteResponse, err := client.DoRequest(s.client, url)
	if err != nil {
		return
	}

	var schedule []Schedule
	err = json.Unmarshal(byteResponse, &schedule)
	if err != nil {
		return
	}
	
	return
}