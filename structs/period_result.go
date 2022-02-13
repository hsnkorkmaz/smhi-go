package structs

type PeriodResult struct {
	Key     string `json:"key"`
	Updated int64  `json:"updated"`
	Title   string `json:"title"`
	Summary string `json:"summary"`
	From    int64  `json:"from"`
	To      int64  `json:"to"`
	Links    []Link `json:"link"`
	Datas    []Data `json:"data"`
}
