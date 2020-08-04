package entities

type Anime struct {
	Id int64 `json:"id"`
	Title string `json:"title"`
	EpisodeNo uint `json:"episodeNo"`
}