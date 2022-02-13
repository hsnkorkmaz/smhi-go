package structs

type ParameterResult struct {
	Key        string       `json:"key"`
	Updated    int64        `json:"updated"`
	Title      string       `json:"title"`
	Summary    string       `json:"summary"`
	ValueType  string       `json:"valueType"`
	Links       []Link       `json:"link"`
	StationSets []StationSet `json:"stationSet"`
	Stations    []Station    `json:"station"`
}