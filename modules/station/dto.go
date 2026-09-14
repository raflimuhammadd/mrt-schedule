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