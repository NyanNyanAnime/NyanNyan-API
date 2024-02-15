package service

import (
	"errors"
	"mime/multipart"
	"nyannyan/features/anime/entity"
	"nyannyan/utils/constanta"
	"nyannyan/utils/validation"
)

type animeService struct {
	animeRepository entity.AnimeRepositoryInterface
}

func NewAnimeService(anime entity.AnimeRepositoryInterface) entity.AnimeServiceInterface {
	return &animeService{
		animeRepository: anime,
	}
}

// CreateAnime implements entity.AnimeServiceInterface.
func (as *animeService) CreateAnime(image *multipart.FileHeader, data entity.AnimeCore) error {
	errEmpty := validation.CheckDataEmpty(data.Title, data.Synopsis)
	if errEmpty != nil {
		return errEmpty
	}

	_, err := as.animeRepository.GetAnimeByTitle(data.Title)
	if err == nil {
		return errors.New("title is already in use")
	}

	if image.Size > 10*1024*1024 {
		return errors.New("image file size should be less than 10 MB")
	}

	errCreate := as.animeRepository.CreateAnime(image, data)
	if errCreate != nil {
		return errCreate
	}

	return nil
}

// DeleteAnimeById implements entity.AnimeServiceInterface.
func (as*animeService) DeleteAnimeById(id string) error {
	if id == "" {
		return errors.New(constanta.ERROR_ID_INVALID)
	}

	err := as.animeRepository.DeleteAnimeById(id)
	if err != nil {
		return err
	}

	return nil
}

// GetAllAnime implements entity.AnimeServiceInterface.
func (as *animeService) GetAllAnime() ([]entity.AnimeCore, error) {
	anime, err := as.animeRepository.GetAllAnime()
	if err != nil {
		return nil, err
	}

	return anime, nil
}

// GetAnimeById implements entity.AnimeServiceInterface.
func (as *animeService) GetAnimeById(id string) (entity.AnimeCore, error) {
	if id == "" {
		return entity.AnimeCore{}, errors.New(constanta.ERROR_ID_INVALID)
	}

	idAnime, err := as.animeRepository.GetAnimeById(id)
	if err != nil {
		return entity.AnimeCore{}, err
	}

	return idAnime, err
}

// UpdateAnimeById implements entity.AnimeServiceInterface.
func (as *animeService) UpdateAnimeById(id string, image *multipart.FileHeader, updated entity.AnimeCore) error {
	errEmpty := validation.CheckDataEmpty(updated.Title, updated.Synopsis)
	if errEmpty != nil {
		return errEmpty
	}

	_, err := as.animeRepository.GetAnimeByTitle(updated.Title)
	if err == nil {
		return errors.New("title is already in use")
	}

	if image.Size > 10*1024*1024 {
		return errors.New("image file size should be less than 10 MB")
	}

	err = as.animeRepository.UpdateAnimeById(id, image, updated)
	if err != nil {
		return err
	}

	return nil
}
