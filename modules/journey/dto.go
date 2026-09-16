package journey

type JourneyResponse struct {
	From        StationDetail `json:"from"`
	To          StationDetail `json:"to"`
	Stations    []StationInfo `json:"stations"`
	StationsCount int         `json:"stations_count"`
	Fare        int           `json:"fare"`
	Duration    int           `json:"duration_minutes"`
	Direction   string        `json:"direction"`
}

type StationDetail struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Slug string `json:"slug"`
}

type StationInfo struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Order int    `json:"order"`
}

type FareResponse struct {
	From       StationDetail `json:"from"`
	To         StationDetail `json:"to"`
	Fare       int           `json:"fare"`
	Stations   int           `json:"stations_count"`
}

type DurationResponse struct {
	From       StationDetail `json:"from"`
	To         StationDetail `json:"to"`
	Duration   int           `json:"duration_minutes"`
	Stations   int           `json:"stations_count"`
}