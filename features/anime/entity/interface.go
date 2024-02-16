package entity

import (
	"mime/multipart"
	"nyannyan/utils/pagination"
)

type AnimeRepositoryInterface interface {
	CreateAnime(image *multipart.FileHeader, data AnimeCore) error
	GetAllAnime(page, limit int, search string) ([]AnimeCore, pagination.PageInfo, int, error)
	GetAnimeById(id string) (AnimeCore, error)
	GetAnimeByTitle(title string) (AnimeCore, error)
	UpdateAnimeById(id string, image *multipart.FileHeader, updated AnimeCore) error
	DeleteAnimeById(id string) error
	CreateGenre(data []GenreCore) error
	UpdateGenreById(id string, data GenreCore) error
	DeleteGenreById(id string) error
}

type AnimeServiceInterface interface {
	CreateAnime(image *multipart.FileHeader, data AnimeCore) error
	GetAllAnime(page, limit int, search string) ([]AnimeCore, pagination.PageInfo, int, error)
	GetAnimeById(id string) (AnimeCore, error)
	UpdateAnimeById(id string, image *multipart.FileHeader, updated AnimeCore) error
	DeleteAnimeById(id string) error
	CreateGenre(data []GenreCore) error
	UpdateGenreById(id string, data GenreCore) error
	DeleteGenreById(id string) error
}
