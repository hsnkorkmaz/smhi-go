package structs

type Period struct {
	Key      string `json:"key"`
	From     int64  `json:"from"`
	To       int64  `json:"to"`
	Updated  int64  `json:"updated"`
	Title    string `json:"title"`
	Summary  string `json:"summary"`
	Sampling string `json:"sampling"`
	Links    []Link `json:"link"`
}
