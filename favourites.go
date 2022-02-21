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

//go:generate mockery --name FavouritesSaver
type FavouriteSaver interface {
	Save(Favourite) error
}

//go:generate mockery --name FavouritesDestroyer
type FavouriteDestroyer interface {
	Destroy(FilmID, UserID) error
}
