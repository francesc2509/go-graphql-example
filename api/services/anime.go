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
	return animes, nil
}

// Create an anime
func (service *animeService) Create(anime *entities.Anime) error {
	anime.Id = uint64(time.Now().UnixNano() / int64(time.Millisecond))

	animes = append(animes, anime)

	return nil
}

func (service *animeService) Modify(anime *entities.Anime) (*entities.Anime, error) {
	idx, err := findIndex(anime.Id)

	if err != nil {
		return nil, err
	}

	item := animes[idx]

	item.Title = anime.Title
	item.EpisodeNo = anime.EpisodeNo

	return item, nil
}

func (service *animeService) Delete(id uint64) (*entities.Anime, error) {
	idx, err := findIndex(id)

	if err != nil {
		return nil, err
	}

	item := animes[idx]
	animes = append(animes[:idx], animes[idx+1:]...)

	return item, nil
}

func findIndex(id uint64) (int, error) {
	for pos, item := range animes {
		if item.Id == id {
			return pos, nil
		}
	}
	return -1, errors.New("The anime was not found")
}
