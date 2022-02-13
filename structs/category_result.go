package structs

type CategoryResult struct {
	Key      string    `json:"key"`
	Updated  int64     `json:"updated"`
	Title    string    `json:"title"`
	Summary  string    `json:"summary"`
	Links    []Link    `json:"link"`
	Versions []Version `json:"version"`
}
