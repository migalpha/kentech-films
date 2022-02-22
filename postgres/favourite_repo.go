package postgres

import (
	"context"
	"fmt"

	film "github.com/migalpha/kentech-films"
	"gorm.io/gorm"
)

type FavouriteRepo struct {
	DB *gorm.DB
}

func NewFavouriteRepository(db *gorm.DB) *FavouriteRepo {
	return &FavouriteRepo{DB: db}
}

func (repo FavouriteRepo) Save(ctx context.Context, data film.Favourite) (film.FavouriteID, error) {
	favouriteData := encodeFavourite(data)
	result := repo.DB.Create(&favouriteData)
	if result.Error != nil {
		return 0, fmt.Errorf("[FavouriteRepo:Save][err:%w]", result.Error)
	}
	return film.FavouriteID(favouriteData.ID), nil
}

func (repo FavouriteRepo) Destroy(ctx context.Context, filmID film.FilmID, userID film.UserID) error {
	resp := repo.DB.Where("film_id = ? AND user_id = ?", filmID, userID).Delete(&PostgresFavourite{})
	return resp.Error
}

type PostgresFavourite struct {
	ID     uint `json:"id"`
	FilmID uint `json:"film_id"`
	UserID uint `json:"user_id"`
}

func encodeFavourite(favourite film.Favourite) PostgresFavourite {
	return PostgresFavourite{
		FilmID: favourite.FilmID.Uint(),
		UserID: favourite.UserID.Uint(),
	}
}

// TableName overrides the table name used by PostgresFavourite to `favourites`
func (PostgresFavourite) TableName() string {
	return "favourites"
}
