package structs

type Value struct {
	Date    int64  `json:"date"`
	Value   string `json:"value"`
	Quality string `json:"quality"`
}
