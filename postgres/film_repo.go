package postgres

import (
	"context"
	"fmt"
	"log"

	film "github.com/migalpha/kentech-films"

	"gorm.io/gorm"
)

type FilmRepo struct {
	DB *gorm.DB
}

func NewFilmRepository(db *gorm.DB) *FilmRepo {
	return &FilmRepo{DB: db}
}
func (repo FilmRepo) Save(ctx context.Context, data film.Film) (film.FilmID, error) {
	filmData := encodeFilm(data)
	result := repo.DB.Create(&filmData)
	if result.Error != nil {
		return 0, fmt.Errorf("[FilmRepo:Save][err:%w]", result.Error)
	}
	return film.FilmID(filmData.ID), nil
}

func (repo FilmRepo) Destroy(ctx context.Context, ID film.FilmID) error {
	resp := repo.DB.Delete(&PostgresFilm{}, ID.Uint())
	return resp.Error
}

func (repo FilmRepo) FilmbyID(ctx context.Context, ID film.FilmID) (film.Film, error) {
	data := PostgresFilm{
		ID: ID.Uint(),
	}

	res := repo.DB.Where(&data).First(&data)
	if res.Error != nil {
		return film.Film{}, fmt.Errorf("[FilmRepo:FilmbyId][err:%w]", res.Error)
	}
	filmDB, err := repo.decodeFilm(data)
	if err != nil {
		return film.Film{}, fmt.Errorf("can't decode user: %w", err)
	}
	return filmDB, nil
}

func (repo FilmRepo) GetFilms(ctx context.Context, filter map[string]interface{}) (film.Films, error) {
	var filmsDB []PostgresFilm

	res := repo.DB.Where(filter).Find(&filmsDB)
	if res.Error != nil {
		return film.Films{}, fmt.Errorf("[FilmRepo:GetFilms][err:%w]", res.Error)
	}

	var films film.Films
	films.Films = []film.Film{}
	for _, f := range filmsDB {
		decodeFilm, err := repo.decodeFilm(f)
		if err != nil {
			log.Printf("[FilmRepo:GetFilms][decodeFilm][err:%s]", err.Error())
		}
		films.Films = append(films.Films, decodeFilm)
	}

	return films, nil
}

func (repo FilmRepo) Update(ctx context.Context, data film.Film) error {
	fmt.Printf("%+v", data)
	res := repo.DB.Model(&PostgresFilm{}).Where("id = ?", data.ID).Updates(encodeFilm(data))
	return res.Error
}

func (repo FilmRepo) decodeFilm(data PostgresFilm) (film.Film, error) {
	return film.Film{
		ID:          film.FilmID(data.ID),
		Starring:    data.Starring,
		Director:    data.Director,
		Genre:       data.Genre,
		Sypnosis:    data.Sypnosis,
		Title:       data.Title,
		ReleaseYear: data.ReleaseYear,
		CreatedBy:   data.CreatedBy,
	}, nil
}

type PostgresFilm struct {
	ID          uint   `json:"id"`
	Starring    string `json:"starring"`
	Director    string `json:"director"`
	Genre       string `json:"genre"`
	Sypnosis    string `json:"sypnosis"`
	Title       string `json:"title"`
	ReleaseYear int    `json:"release_date"`
	CreatedBy   uint   `json:"created_by"`
}

func encodeFilm(film film.Film) PostgresFilm {
	return PostgresFilm{
		Starring:    film.Starring,
		Director:    film.Director,
		Genre:       film.Genre,
		Sypnosis:    film.Sypnosis,
		Title:       film.Title,
		ReleaseYear: film.ReleaseYear,
		CreatedBy:   film.CreatedBy,
	}
}

type Tabler interface {
	TableName() string
}

// TableName overrides the table name used by PostgresFilm to `films`
func (PostgresFilm) TableName() string {
	return "films"
}
