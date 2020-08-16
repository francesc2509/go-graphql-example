package services

import (
	"errors"
	"time"

	"github.com/francesc2509/go-graphql-example/entities"
)

var animes = []*entities.Anime{
	&entities.Anime{
		Title:     "TTGL",
		EpisodeNo: 27,
	},
}

type animeService struct{}

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

	return animes, nil
}

// Create an anime
func (service *animeService) Create(anime *entities.Anime) error {
	anime.Id = uint64(time.Now().UnixNano() / int64(time.Millisecond))

	animes = append(animes, anime)

	return nil
}

func (service *animeService) Modify(anime *entities.Anime) (*entities.Anime, error) {
	// err := repositories.Anime.Create(*anime)

	item, err := findOne(anime)

	if err != nil {
		return nil, err
	}

	item.Title = anime.Title
	item.EpisodeNo = anime.EpisodeNo

	return item, nil
}

func findOne(anime *entities.Anime) (*entities.Anime, error) {
	for _, item := range animes {
		if item.Id == anime.Id {
			return item, nil
		}
	}
	return nil, errors.New("The anime was not found")
}
