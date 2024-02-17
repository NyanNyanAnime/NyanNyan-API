package repository

import (
	"errors"
	"mime/multipart"
	"nyannyan/features/anime/entity"
	"nyannyan/features/anime/model"
	"nyannyan/utils/constanta"
	"nyannyan/utils/pagination"
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

// CreateGenre implements entity.AnimeRepositoryInterface.
func (ar *animeRepository) CreateGenre(data []entity.GenreCore) error {
	request := entity.ListMapCoreGenreToModelGenre(data)

	tx := ar.db.Create(&request)
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}

// DeleteGenreById implements entity.AnimeRepositoryInterface.
func (ar *animeRepository) DeleteGenreById(id string) error {
	genreData := model.Genre{}

	tx := ar.db.Where("id = ?", id).Delete(&genreData)
	if tx.Error != nil {
		return tx.Error
	}

	if tx.RowsAffected == 0 {
		return errors.New(constanta.ERROR_DATA_NOT_FOUND)
	}

	return nil
}

// UpdateGenreById implements entity.AnimeRepositoryInterface.
func (ar *animeRepository) UpdateGenreById(id string, data entity.GenreCore) error {
	request := entity.MapCoreGenreToModelGenre(data)

	tx := ar.db.Where("id = ?", id).Updates(&request)
	if tx.Error != nil {
		return tx.Error
	}

	if tx.RowsAffected == 0 {
		return errors.New(constanta.ERROR_DATA_NOT_FOUND)
	}

	return nil
}

// CreateAnime implements entity.AnimeRepositoryInterface.
func (ar *animeRepository) CreateAnime(image *multipart.FileHeader, data entity.AnimeCore) error {
	request := entity.MapCoreAnimeToModelAnime(data)

	// Upload image to Cloudinary
	file, err := image.Open()
	if err != nil {
		return err
	}
	defer file.Close()

	imageURL, err := storage.UploadToCloudinary(file, image.Filename)
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
func (ar *animeRepository) GetAllAnime(page, limit int, search string) ([]entity.AnimeCore, pagination.PageInfo, int, error) {
	dataAnime := []model.Anime{}

	offset := (page - 1) * limit
	query := ar.db.Model(&model.Anime{}).Preload("Genre")

	if search != "" {
		query = query.
			Joins("JOIN genres ON genres.anime_id = animes.id").
			Where("animes.title LIKE ? OR genres.genre LIKE ?", "%"+search+"%", "%"+search+"%")
	}

	var totalCount int64
	tx := query.Count(&totalCount).Find(&dataAnime)
	if tx.Error != nil {
		return nil, pagination.PageInfo{}, 0, tx.Error
	}

	query = query.Offset(offset).Limit(limit).Group("animes.id")
	tx = query.Find(&dataAnime)
	if tx.Error != nil {
		return nil, pagination.PageInfo{}, 0, tx.Error
	}

	dataResponse := entity.ListModelAnimeToCoreAnime(dataAnime)
	pageInfo := pagination.CalculateData(int(totalCount), limit, page)

	return dataResponse, pageInfo, int(totalCount), nil
}


// GetAnimeById implements entity.AnimeRepositoryInterface.
func (ar *animeRepository) GetAnimeById(id string) (entity.AnimeCore, error) {
	dataAnime := model.Anime{}

	tx := ar.db.Preload("Genre").Where("id = ?", id).First(&dataAnime)
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
		file, err := image.Open()
		if err != nil {
			return err
		}
		defer file.Close()

		imageURL, uploadErr := storage.UploadToCloudinary(file, image.Filename)
		if uploadErr != nil {
			return uploadErr
		}
		request.Image = imageURL
	}

	// Only update the image if the `image` parameter is not nil
	if image != nil {
		tx = ar.db.Where("id = ?", id).Updates(&request)
		if tx.Error != nil {
			return tx.Error
		}

		if tx.RowsAffected == 0 {
			return errors.New(constanta.ERROR_DATA_NOT_FOUND)
		}
	} else {
		tx = ar.db.Where("id = ?", id).Omit("image").Updates(&request)
		if tx.Error != nil {
			return tx.Error
		}

		if tx.RowsAffected == 0 {
			return errors.New(constanta.ERROR_DATA_NOT_FOUND)
		}
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
