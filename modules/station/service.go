package station

import (
	"encoding/json"
	"errors"
	"net/http"
	"os"
	"strings"
	"time"
	"net/url"

	"github.com/raflimuhammadd/mrt-schedule/common/client"
)

type Service interface {
	GetAllStation() (response []StationResponse, err error)
	CheckScheduleByStation(id string) (response []StationResponse, err error)
	GetStationScheduleBySlug(slug string) (response StationScheduleResponse, err error)
}

type service struct {
	client *http.Client
	baseURL string
}

func isWeekday(time.Time) bool {
	today := time.Now().Weekday()
	return today != time.Sunday && today != time.Saturday
}

func parseScheduleTime(timeString string) []string {
	if timeString == "" {
		return []string{}
	}
	times := strings.Split(timeString, "; ")
	return times
}

func getDayType(t time.Time) string {
	if isWeekday(t) {
		return "weekday"
	}
	return "weekend"
}

func buildStationScheduleURL(baseURL, slug string) string {
	params := url.Values{}

	params.Add("fields[]", "id")
	params.Add("fields[]", "name")
	params.Add("fields[]", "slug")
	params.Add("fields[]", "object")

	params.Add("filters[field][slug]", "stasiun")
	params.Add("filters[slug]", slug)

	params.Add("pagination[limit]", "1")
	params.Add("sort[]", "id:desc")

	return baseURL + "/datum?" + params.Encode()
}

func formatScheduleResponse(stationID int, stationName string, startTimes []string, endTimes []string) StationScheduleResponse {
    now := time.Now() // Current time
    
    // Filter both directions
    filteredStartTimes := filterFutureTimes(startTimes, now)
    filteredEndTimes := filterFutureTimes(endTimes, now)
    
    // Rest of the function stays the same
    return StationScheduleResponse{
        StationID:   stationID,
        StationName: stationName,
        DayType:     getDayType(now),
        Directions: []DirectionSchedule{
            {
                Direction:   "to_bundaran_hi",
                Destination: "Bundaran HI",
                Times:       filteredEndTimes,  // Filtered
            },
            {
                Direction:   "to_lebak_bulus",
                Destination: "Lebak Bulus",
                Times:       filteredStartTimes, // Filtered
            },
        },
    }
}

func filterFutureTimes(times []string, currentTime time.Time) []string {
    var futureTimes []string
    
    for _, timeStr := range times {
        // Parse time string "HH:MM:SS" to time.Time
        scheduleTime, err := time.Parse("15:04:05", timeStr)
        if err != nil {
            // Skip invalid time format
            continue
        }
        
        // Combine current date with schedule time
        scheduleDateTime := time.Date(
            currentTime.Year(),
            currentTime.Month(),
            currentTime.Day(),
            scheduleTime.Hour(),
            scheduleTime.Minute(),
            scheduleTime.Second(),
            0, // nanoseconds
            currentTime.Location(),
        )
        
        // Compare: is this schedule in the future?
        if scheduleDateTime.After(currentTime) {
            futureTimes = append(futureTimes, timeStr)
        }
    }
    
    return futureTimes
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

func (s *service) GetStationScheduleBySlug(slug string) (response StationScheduleResponse, err error) {
    baseURL := os.Getenv("API_BASE_URL")
    if baseURL == "" {
        return response, errors.New("API_BASE_URL is required")
    }
    
    // Build URL with query parameters
    url := buildStationScheduleURL(baseURL, slug)
    
    // Make API request
    byteResponse, err := client.DoRequest(s.client, url)
    if err != nil {
        return response, err
    }
    
    // Parse response (API returns array even with limit=1)
    var apiResp struct {
        Data []StationData `json:"data"`
    }
    err = json.Unmarshal(byteResponse, &apiResp)
    if err != nil {
        return response, err
    }
    
    // Validate: check if station found
    if len(apiResp.Data) == 0 {
        return response, errors.New("station not found")
    }
    
    stationData := apiResp.Data[0]  // Get first (and only) result
    
    // Validate schedule data exists
    if stationData.Object.Schedule.WeekdaysStart == "" {
        return response, errors.New("schedule data not available for this station")
    }
    
    // Determine day type and parse schedules
    var startTimes, endTimes []string
    if isWeekday(time.Now()) {
        startTimes = parseScheduleTime(stationData.Object.Schedule.WeekdaysStart)
        endTimes = parseScheduleTime(stationData.Object.Schedule.WeekdaysEnd)
    } else {
        startTimes = parseScheduleTime(stationData.Object.Schedule.WeekendsStart)
        endTimes = parseScheduleTime(stationData.Object.Schedule.WeekendsEnd)
    }
    
    // Format response
    response = formatScheduleResponse(
        stationData.ID,
        stationData.Name,
        startTimes,
        endTimes,
    )
    
    return response, nil
}

func (s *service) CheckScheduleByStation(id string) (response []StationResponse, err error) {
	baseURL := os.Getenv("API_BASE_URL")
	if baseURL == "" {
		return nil, errors.New("API_BASE_URL is required")
	}

	url := baseURL + "/datum/" + id

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