package film

import "context"

type Favourite struct {
	ID     FavouriteID `json:"id"`
	FilmID FilmID      `json:"film_id"`
	UserID UserID      `json:"user_id"`
}

type FavouriteID uint

func (id FavouriteID) Uint() uint {
	return uint(id)
}

//go:generate mockery --name FavouriteSaver
type FavouriteSaver interface {
	Save(context.Context, Favourite) (FavouriteID, error)
}

//go:generate mockery --name FavouriteDestroyer
type FavouriteDestroyer interface {
	Destroy(context.Context, FilmID, UserID) error
}
