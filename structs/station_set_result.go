package structs

type StationSetResult struct {
	Key           string     `json:"key"`
	Updated       int64      `json:"updated"`
	Title         string     `json:"title"`
	Owner         string     `json:"owner"`
	OwnerCategory string     `json:"ownerCategory"`
	Active        bool       `json:"active"`
	Summary       string     `json:"summary"`
	From          int64      `json:"from"`
	To            int64      `json:"to"`
	Positions     []Position `json:"position"`
	Links         []Link     `json:"link"`
	Periods       []Period   `json:"period"`
}
