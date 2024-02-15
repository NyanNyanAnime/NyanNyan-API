package entity

import "mime/multipart"

type AnimeRepositoryInterface interface {
	CreateAnime(image *multipart.FileHeader, data AnimeCore) error
	GetAllAnime() ([]AnimeCore, error)
	GetAnimeById(id string) (AnimeCore, error)
	GetAnimeByTitle(title string) (AnimeCore, error)
	UpdateAnimeById(id string, image *multipart.FileHeader, updated AnimeCore) error
	DeleteAnimeById(id string) error
}

type AnimeServiceInterface interface {
	CreateAnime(image *multipart.FileHeader, data AnimeCore) error
	GetAllAnime() ([]AnimeCore, error)
	GetAnimeById(id string) (AnimeCore, error)
	UpdateAnimeById(id string, image *multipart.FileHeader, updated AnimeCore) error
	DeleteAnimeById(id string) error
}
