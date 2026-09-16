package journey

import (
	"errors"
)

// Service interface untuk journey module
type Service interface {
	PlanJourney(fromSlug, toSlug string) (JourneyResponse, error)
	CalculateFare(fromSlug, toSlug string) (FareResponse, error)
	CalculateDuration(fromSlug, toSlug string) (DurationResponse, error)
}

type service struct {
	// No external dependencies needed for calculation
}

func NewService() Service {
	return &service{}
}

func (s *service) PlanJourney(fromSlug, toSlug string) (response JourneyResponse, err error) {
	// 1. Find from station
	fromStation, err := FindStationBySlug(fromSlug)
	if err != nil {
		return response, err
	}

	// 2. Find to station
	toStation, err := FindStationBySlug(toSlug)
	if err != nil {
		return response, err
	}

	// 3. Validate: same station?
	if fromStation.ID == toStation.ID {
		return response, errors.New("from and to stations cannot be the same")
	}

	// 4. Get station list between from and to
	var stations []StationInfo
	stationsCount := 0

	if fromStation.Order < toStation.Order {
		// Northbound: fromOrder → toOrder
		for i := fromStation.Order - 1; i < toStation.Order; i++ {
			stations = append(stations, StationInfo{
				ID:    StationOrder[i].ID,
				Name:  StationOrder[i].Name,
				Order: StationOrder[i].Order,
			})
			stationsCount++
		}
	} else {
		// Southbound: toOrder → fromOrder
		for i := toStation.Order - 1; i < fromStation.Order; i++ {
			stations = append(stations, StationInfo{
				ID:    StationOrder[i].ID,
				Name:  StationOrder[i].Name,
				Order: StationOrder[i].Order,
			})
			stationsCount++
		}
	}

	// 5. Calculate fare
	fare := CalculateFare(stationsCount)

	// 6. Calculate duration
	duration := CalculateDuration(stationsCount)

	// 7. Get direction
	direction := GetDirection(fromStation.Order, toStation.Order)

	// 8. Build response
	response = JourneyResponse{
		From: StationDetail{
			ID:   fromStation.ID,
			Name: fromStation.Name,
			Slug: fromStation.Slug,
		},
		To: StationDetail{
			ID:   toStation.ID,
			Name: toStation.Name,
			Slug: toStation.Slug,
		},
		Stations:    stations,
		StationsCount: stationsCount,
		Fare:        fare,
		Duration:    duration,
		Direction:   direction,
	}

	return response, nil
}

func (s *service) CalculateFare(fromSlug, toSlug string) (response FareResponse, err error) {
	journey, err := s.PlanJourney(fromSlug, toSlug)
	if err != nil {
		return response, err
	}

	response = FareResponse{
		From:       journey.From,
		To:         journey.To,
		Fare:       journey.Fare,
		Stations:   journey.StationsCount,
	}

	return response, nil
}

func (s *service) CalculateDuration(fromSlug, toSlug string) (response DurationResponse, err error) {
	journey, err := s.PlanJourney(fromSlug, toSlug)
	if err != nil {
		return response, err
	}

	response = DurationResponse{
		From:     journey.From,
		To:       journey.To,
		Duration: journey.Duration,
		Stations: journey.StationsCount,
	}

	return response, nil
}

