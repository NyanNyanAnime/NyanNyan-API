package repository

import (
	"errors"
	"mime/multipart"
	"nyannyan/features/anime/entity"
	"nyannyan/features/anime/model"
	"nyannyan/utils/constanta"
	"nyannyan/utils/storage"

	"gorm.io/gorm"
)

type animeRepository struct {
	db *gorm.DB
}

func NewAnimeRepository(db *gorm.DB) entity.AnimeRepositoryInterface {
	return &animeRepository{
		db: db,
	}
}

// CreateAnime implements entity.AnimeRepositoryInterface.
func (ar *animeRepository) CreateAnime(image *multipart.FileHeader, data entity.AnimeCore) error {
	request := entity.MapCoreAnimeToModelAnime(data)

	imageURL, err := storage.UploadImage(image)
	if err != nil {
		return err
	}
	request.Image = imageURL

	tx := ar.db.Create(&request)
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}

// DeleteAnimeById implements entity.AnimeRepositoryInterface.
func (ar *animeRepository) DeleteAnimeById(id string) error {
	animeData := model.Anime{}

	tx := ar.db.Unscoped().Where("id = ?", id).Delete(&animeData)
	if tx.Error != nil {
		return tx.Error
	}

	if tx.RowsAffected == 0 {
		return errors.New(constanta.ERROR_DATA_NOT_FOUND)
	}

	return nil
}

// GetAllAnime implements entity.AnimeRepositoryInterface.
func (ar *animeRepository) GetAllAnime() ([]entity.AnimeCore, error) {
	dataAnime := []model.Anime{}

	err := ar.db.Find(&dataAnime).Error
	if err != nil {
		return nil, err
	}

	dataResponse := entity.ListModelAnimeToCoreAnime(dataAnime)
	return dataResponse, nil
}

// GetAnimeById implements entity.AnimeRepositoryInterface.
func (ar *animeRepository) GetAnimeById(id string) (entity.AnimeCore, error) {
	dataAnime := model.Anime{}

	tx := ar.db.Preload("Schedule").Where("id = ?", id).First(&dataAnime)
	if tx.Error != nil {
		return entity.AnimeCore{}, tx.Error
	}

	if tx.RowsAffected == 0 {
		return entity.AnimeCore{}, errors.New(constanta.ERROR_DATA_NOT_FOUND)
	}

	dataResponse := entity.ModelAnimeToCoreAnime(dataAnime)
	return dataResponse, nil
}

// UpdateAnimeById implements entity.AnimeRepositoryInterface.
func (ar *animeRepository) UpdateAnimeById(id string, image *multipart.FileHeader, updated entity.AnimeCore) error {
	dataAnime := model.Anime{}

	request := entity.MapCoreAnimeToModelAnime(updated)

	tx := ar.db.Where("id = ?", id).First(&dataAnime)
	if tx.Error != nil {
		return tx.Error
	}

	if image != nil {
		imageURL, uploadErr := storage.UploadImage(image)
		if uploadErr != nil {
			return uploadErr
		}
		request.Image = imageURL
	} else {
		request.Image = dataAnime.Image
	}

	tx = ar.db.Where("id = ?", id).Updates(&request)
	if tx.Error != nil {
		return tx.Error
	}

	if tx.RowsAffected == 0 {
		return errors.New(constanta.ERROR_DATA_NOT_FOUND)
	}

	return nil
}

// GetAnimeByTitle implements entity.AnimeRepositoryInterface.
func (ar *animeRepository) GetAnimeByTitle(title string) (entity.AnimeCore, error) {
	dataAnime := model.Anime{}
	tx := ar.db.Where("name = ?", title).First(&dataAnime)

	if tx.RowsAffected == 0 {
		return entity.AnimeCore{}, errors.New(constanta.ERROR_DATA_NOT_FOUND)
	}

	if tx.Error != nil {
		return entity.AnimeCore{}, tx.Error
	}

	result := entity.ModelAnimeToCoreAnime(dataAnime)
	return result, nil
}
