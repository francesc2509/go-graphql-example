package entities

type Anime struct {
	Id        uint64 `json:"id"`
	Title     string `json:"title"`
	EpisodeNo uint   `json:"episodeNo"`
}
