package film

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
	Save(Favourite) (FavouriteID, error)
}

//go:generate mockery --name FavouriteDestroyer
type FavouriteDestroyer interface {
	Destroy(FilmID, UserID) error
}
