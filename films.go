package film

import (
	"context"
	"fmt"
	"strconv"
)

type Film struct {
	ID          FilmID `json:"id" example:"1"`
	Starring    string `json:"starring" example:"Brad Pitt, Christoph Waltz, Michael Fassbender"`
	Director    string `json:"director" example:"Quentin Tarantino"`
	Genre       string `json:"genre" example:"action"`
	Sypnosis    string `json:"sypnosis"`
	Title       string `json:"title" example:"Inglourious Basterds"`
	ReleaseYear int    `json:"release_year" example:"2009"`
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
	FilmbyID(context.Context, FilmID) (Film, error)
	GetFilms(context.Context, map[string]interface{}) (Films, error)
}

//go:generate mockery --name FilmSaver
type FilmSaver interface {
	Save(context.Context, Film) (FilmID, error)
}

//go:generate mockery --name FilmUpdater
type FilmUpdater interface {
	Update(context.Context, Film) error
}

//go:generate mockery --name FilmDestroyer
type FilmDestroyer interface {
	Destroy(context.Context, FilmID) error
}
