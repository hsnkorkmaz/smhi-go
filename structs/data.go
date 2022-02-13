package structs

type Data struct {
	Key     interface{} `json:"key"`
	Updated int64       `json:"updated"`
	Title   string      `json:"title"`
	Summary string      `json:"summary"`
	Links    []Link      `json:"link"`
}