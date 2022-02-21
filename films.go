package film

import (
	"fmt"
	"strconv"
)

type Film struct {
	ID          FilmID `json:"id"`
	Starring    string `json:"starring"`
	Director    string `json:"director"`
	Genre       string `json:"genre"`
	Sypnosis    string `json:"sypnosis"`
	Title       string `json:"title"`
	ReleaseYear int    `json:"release_year"`
	CreatedBy   uint   `json:"created_by"`
}

type Films struct {
	Films []Film `json:"films"`
}

type FilmID uint

func (id FilmID) Uint() uint {
	return uint(id)
}

func NewFilmID(id string) (FilmID, error) {
	value, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		return 0, fmt.Errorf("%v: %w", err, ErrInvalidFilmID)
	}
	return FilmID(value), nil
}

//go:generate mockery --name FilmProvider
type FilmProvider interface {
	FilmbyID(FilmID) (Film, error)
	GetFilms(map[string]interface{}) (Films, error)
}

//go:generate mockery --name FilmSaver
type FilmSaver interface {
	Save(Film) error
}

//go:generate mockery --name FilmUpdater
type FilmUpdater interface {
	Update(Film) error
}

//go:generate mockery --name FilmDestroyer
type FilmDestroyer interface {
	Destroy(FilmID) error
}
