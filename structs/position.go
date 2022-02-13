package structs

type Position struct {
	From      int64   `json:"from"`
	To        int64   `json:"to"`
	Height    float64 `json:"height"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}
