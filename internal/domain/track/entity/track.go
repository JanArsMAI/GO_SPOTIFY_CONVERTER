package entity

type Track struct {
	Name       string `json:"name"`
	ID         string `json:"id"`
	DurationMs int    `json:"duration_ms"`
	Popularity int    `json:"popularity"`
	PreviewURL string `json:"preview_url"`
	Album      struct {
		Name        string `json:"name"`
		ReleaseDate string `json:"release_date"`
	} `json:"album"`
	Artists []struct {
		Name string `json:"name"`
		ID   string `json:"id"`
	} `json:"artists"`
}
