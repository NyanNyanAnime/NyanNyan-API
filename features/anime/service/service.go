package service

import (
	"errors"
	"mime/multipart"
	"nyannyan/features/anime/entity"
	"nyannyan/utils/constanta"
	"nyannyan/utils/pagination"
	"nyannyan/utils/validation"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

type animeService struct {
	animeRepository entity.AnimeRepositoryInterface
}

func NewAnimeService(anime entity.AnimeRepositoryInterface) entity.AnimeServiceInterface {
	return &animeService{
		animeRepository: anime,
	}
}

// CreateGenre implements entity.AnimeServiceInterface.
func (as *animeService) CreateGenre(data []entity.GenreCore) error {
	var updatedGenres []entity.GenreCore

	for _, genre := range data {
		if errEmpty := validation.CheckDataEmpty(genre.Genre, genre.AnimeId); errEmpty != nil {
			return errEmpty
		}

		tc := cases.Title(language.English)
		genre.Genre = tc.String(genre.Genre)

		updatedGenres = append(updatedGenres, genre)
	}

	errCreate := as.animeRepository.CreateGenre(updatedGenres)
	if errCreate != nil {
		return errCreate
	}

	return nil
}

// DeleteGenreById implements entity.AnimeServiceInterface.
func (as *animeService) DeleteGenreById(id string) error {
	if id == "" {
		return errors.New(constanta.ERROR_ID_INVALID)
	}

	err := as.animeRepository.DeleteGenreById(id)
	if err != nil {
		return err
	}

	return nil
}

// UpdateGenreById implements entity.AnimeServiceInterface.
func (as *animeService) UpdateGenreById(id string, data entity.GenreCore) error {
	if id == "" {
		return errors.New(constanta.ERROR_ID_INVALID)
	}

	errEmpty := validation.CheckDataEmpty(data.Genre)
	if errEmpty != nil {
		return errEmpty
	}

	errUpdate := as.animeRepository.UpdateGenreById(id, data)
	if errUpdate != nil {
		return errUpdate
	}

	return nil
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

	tc := cases.Title(language.English)
	data.Title = tc.String(data.Title)

	errCreate := as.animeRepository.CreateAnime(image, data)
	if errCreate != nil {
		return errCreate
	}

	return nil
}

// DeleteAnimeById implements entity.AnimeServiceInterface.
func (as *animeService) DeleteAnimeById(id string) error {
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
func (as *animeService) GetAllAnime(page, limit int, search string) ([]entity.AnimeCore, pagination.PageInfo, int, error) {
	if limit > 10 {
		return nil, pagination.PageInfo{}, 0, errors.New("the limit cannot be more than 10")
	}

	page, limit = validation.ValidateCountLimitAndPage(page, limit)

	search = cases.Title(language.English).String(search)

	animeCores, pageInfo, count, err := as.animeRepository.GetAllAnime(page, limit, search)
	if err != nil {
		return nil, pagination.PageInfo{}, 0, err
	}

	return animeCores, pageInfo, count, nil
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
	if id == "" {
		return errors.New(constanta.ERROR_ID_INVALID)
	}

	// errEmpty := validation.CheckDataEmpty(updated.Title, updated.Synopsis)
	// if errEmpty != nil {
	// 	return errEmpty
	// }

	if image != nil && image.Size > 10*1024*1024 {
		return errors.New("image file size should be less than 10 MB")
	}

	err := as.animeRepository.UpdateAnimeById(id, image, updated)
	if err != nil {
		return err
	}

	return nil
}
