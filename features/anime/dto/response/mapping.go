package response

import "nyannyan/features/anime/entity"

func CoreAnimeToAnimeResponse(data entity.AnimeCore) AnimeResponse {
	return AnimeResponse{
		Id:       data.Id,
		Title:    data.Title,
		Synopsis: data.Synopsis,
		Image:    data.Image,
		Genres:   ListCoreGenreToGenreRequest(data.Genre),
	}
}

func CoreGenreToGenreResponse(data entity.GenreCore) GenreResponse {
	return GenreResponse{
		Genre: data.Genre,
	}
}

func ListCoreGenreToGenreRequest(data []entity.GenreCore) []GenreResponse {
	list := []GenreResponse{}
	for _, value := range data {
		result := CoreGenreToGenreResponse(value)
		list = append(list, result)
	}
	return list
}

func ListCoreAnimeToAnimeResponse(data []entity.AnimeCore) []AnimeResponse {
	list := []AnimeResponse{}
	for _, value := range data {
		result := CoreAnimeToAnimeResponse(value)
		list = append(list, result)
	}
	return list
}
