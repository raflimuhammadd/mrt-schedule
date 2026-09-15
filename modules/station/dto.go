package station

type Station struct {
	Id int
	Name string
	Description string
}

type StationAPIResponse struct {
	Data []Station `json:"data"`
}

type StationResponse struct {
	Id 				int `json:"id"`
	Name 			string `json:"name"`
	Description 	string `json:"description"`
}

type Schedule struct {
	StationId 			int `json:"id"`
	StationName 		string `json:"name"`
	ScheduleBundaranHI	string `json:"end"`
	ShecduleLebakBulus	string `json:"start"`
}

type ScheduleResponse struct {
	StationName			string `json:"name"`
	Time				string `json:"time"`
}

type StationData struct {
	ID int `json:"id"`
	Name string `json:"name"`
	Object struct {
		Schedule struct {
			WeekdaysStart string `json:"weekdaysStart"`
			WeekdaysEnd string `json:"weekdaysEnd"`
			WeekendsStart string `json:"weekendsStart"`
			WeekendsEnd string `json:"weekendsEnd"`
		} `json:"schedule"`
	} `json:"object"`
}

type StationScheduleResponse struct {
	StationID int `json:"station_id"`
	StationName string `json:"station_name"`
	DayType string `json:"day_type"`
	Directions []DirectionSchedule `json:"directions"`
}

type DirectionSchedule struct {
	Direction string `json:"direction"`
	Destination string `json:"destination"`
	Times []string `json:"times"`
}