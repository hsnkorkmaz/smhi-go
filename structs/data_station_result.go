package structs

type DataStationResult struct {
	Values    []Value    `json:"value"`
	Updated   int64      `json:"updated"`
	Parameter Parameter  `json:"parameter"`
	Station   Station    `json:"station"`
	Period    Period     `json:"period"`
	Positions []Position `json:"position"`
	Link      []Link     `json:"link"`
}
