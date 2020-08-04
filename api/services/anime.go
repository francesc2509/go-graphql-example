package services

import (
	"github.com/francesc2509/go-graphql-example/entities"
	"time"
	"fmt"
)

type animeService struct {}

// Get a slice of anime
func (service *animeService) Get() ([]*entities.Anime, error) {
	//result, err := repositories.anime.Get()

	// if err != nil {
	// 	return nil, err
	// }

	// // var anime []*entities.anime
	// for _, item := range result {
	// 	anime = append(anime, item.(*entities.anime))
	// }

	animes := []*entities.Anime{
		&entities.Anime {
			Title: "TTGL",
			EpisodeNo: 27,
		},
	}

	return animes, nil
}

// Create an anime
func (service *animeService) Create(anime *entities.Anime) (error) {
	// err := repositories.Anime.Create(*anime)

	anime.Id = time.Now().UnixNano() / int64(time.Millisecond)
	fmt.Println(anime.Id)
	return nil
}