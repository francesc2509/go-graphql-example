package services

import (
	"github.com/francesc2509/go-graphql-example/entities"
)

var Anime = &animeService{}

type animeService struct {}

// Get a slice of artists
func (service *animeService) Get() ([]*entities.Anime, error) {
	//result, err := repositories.Artist.Get()

	// if err != nil {
	// 	return nil, err
	// }

	// // var artists []*entities.Artist
	// for _, item := range result {
	// 	artists = append(artists, item.(*entities.Artist))
	// }

	animes := []*entities.Anime{
		&entities.Anime {
			Title: "TTGL",
			EpisodeNo: 27,
		},
	}

	return animes, nil
}

// Create an artist
func (service *animeService) Create(artist *entities.Anime) (*entities.Anime, error) {
	// err := repositories.Artist.Create(*artist)

	return nil, nil
}