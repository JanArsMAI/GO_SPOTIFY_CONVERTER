package entity

type Track struct {
	Name        string `json:"name"`
	ID          string `json:"id"`
	DurationMs  int    `json:"duration_ms"`
	Popularity  int    `json:"popularity"`
	PreviewURL  string `json:"preview_url"`
	Explicit    bool   `json:"explicit"`
	ExternalIDs struct {
		ISRC string `json:"isrc"`
	} `json:"external_ids"`
	Album   Album    `json:"album"`
	Artists []Artist `json:"artists"`
}

type Album struct {
	Name                 string  `json:"name"`
	ID                   string  `json:"id"`
	AlbumType            string  `json:"album_type"`
	TotalTracks          int     `json:"total_tracks"`
	ReleaseDate          string  `json:"release_date"`
	ReleaseDatePrecision string  `json:"release_date_precision"`
	Images               []Image `json:"images"`
	ExternalURLs         struct {
		Spotify string `json:"spotify"`
	} `json:"external_urls"`
}

type Artist struct {
	Name         string `json:"name"`
	ID           string `json:"id"`
	ExternalURLs struct {
		Spotify string `json:"spotify"`
	} `json:"external_urls"`
}

type Image struct {
	URL    string `json:"url"`
	Height int    `json:"height"`
	Width  int    `json:"width"`
}
