package entity

import "nyannyan/features/anime/model"

// Core To Model
func MapCoreAnimeToModelAnime(data AnimeCore) model.Anime {
	return model.Anime{
		Id:       data.Id,
		Title:    data.Title,
		Synopsis: data.Synopsis,
		Image:    data.Image,
		Genre:    ListMapCoreGenreToModelGenre(data.Genre),
	}
}

func ListMapCoreAnimeToModelAnime(data []AnimeCore) []model.Anime {
	list := []model.Anime{}
	for _, value := range data {
		result := MapCoreAnimeToModelAnime(value)
		list = append(list, result)
	}
	return list
}

func MapCoreGenreToModelGenre(data GenreCore) model.Genre {
	return model.Genre{
		Id:      data.Id,
		Genre:   data.Genre,
		GenreId: data.GenreId,
	}
}

func ListMapCoreGenreToModelGenre(data []GenreCore) []model.Genre {
	list := []model.Genre{}
	for _, value := range data {
		result := MapCoreGenreToModelGenre(value)
		list = append(list, result)
	}
	return list
}

// Model to Core
func ModelAnimeToCoreAnime(data model.Anime) AnimeCore {
	return AnimeCore{
		Id:        data.Id,
		Title:     data.Title,
		Synopsis:  data.Synopsis,
		Image:     data.Image,
		Genre:     ListModelGenreToCoreGenre(data.Genre),
		CreatedAt: data.CreatedAt,
		UpdatedAt: data.UpdatedAt,
	}
}

func ModelGenreToCoreGenre(data model.Genre) GenreCore {
	return GenreCore{
		Id:        data.Id,
		Genre:     data.Genre,
		GenreId:   data.GenreId,
		CreatedAt: data.CreatedAt,
		UpdatedAt: data.UpdatedAt,
	}
}

func ListModelAnimeToCoreAnime(data []model.Anime) []AnimeCore {
	list := []AnimeCore{}
	for _, value := range data {
		result := ModelAnimeToCoreAnime(value)
		list = append(list, result)
	}
	return list
}

func ListModelGenreToCoreGenre(data []model.Genre) []GenreCore {
	list := []GenreCore{}
	for _, value := range data {
		result := ModelGenreToCoreGenre(value)
		list = append(list, result)
	}
	return list
}
