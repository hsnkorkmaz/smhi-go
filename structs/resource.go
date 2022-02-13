package structs

type Resource struct {
	GeoBox  GeoBox `json:"geoBox"`
	Key     string `json:"key"`
	Updated int64  `json:"updated"`
	Title   string `json:"title"`
	Summary string `json:"summary"`
	Links   []Link `json:"link"`
}
