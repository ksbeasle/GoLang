package mock

import "github.com/ksbeasle/GoLang/pkg/models"

var mockGame = &models.Game{
	ID:          1,
	Title:       "Halo 3",
	Genre:       "first-person shooter",
	Rating:      9,
	Platform:    "Xbox 360, Xbox One",
	ReleaseDate: "September 25, 2007",
}

type VGModel struct{}

func (vg *VGModel) Insert(title string, genre string, rating int, platform string, releaseDate string) (int, error) {
	return 2, nil
}

func (vg *VGModel) All() ([]*models.Game, error) {
	return []*models.Game{mockGame}, nil
}

func (vg *VGModel) Get(Id int) (*models.Game, error) {
	switch Id:
case 1:
	return mockGame, nil
default:
	return nil, models.ErrNoGameFound
}
