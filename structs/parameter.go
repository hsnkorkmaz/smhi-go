package structs

type Parameter struct {
	Key     string `json:"key"`
	Name    string `json:"name"`
	Summary string `json:"summary"`
	Unit    string `json:"unit"`
}