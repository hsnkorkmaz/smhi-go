package structs

type DataStationSetResult struct {
	Station   []Station `json:"station"`
	Updated   int64     `json:"updated"`
	Parameter Parameter `json:"parameter"`
	Period    Period    `json:"period"`
	Links     []Link    `json:"link"`
}
