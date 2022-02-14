package structs

type Station struct {
	Name          string  `json:"name"`
	Owner         string  `json:"owner"`
	OwnerCategory string  `json:"ownerCategory"`
	Id            int     `json:"id"`
	Height        float64 `json:"height"`
	Latitude      float64 `json:"latitude"`
	Longitude     float64 `json:"longitude"`
	Active        bool    `json:"active"`
	From          int64   `json:"from"`
	To            int64   `json:"to"`
	Key           string  `json:"key"`
	Updated       int64   `json:"updated"`
	Title         string  `json:"title"`
	Summary       string  `json:"summary"`
	Links         []Link  `json:"link"`
	Values         []Value `json:"value"`
}
