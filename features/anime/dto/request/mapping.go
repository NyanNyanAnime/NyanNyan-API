package request

import "nyannyan/features/anime/entity"

func AnimeRequestToCoreAnime(data AnimeRequest) entity.AnimeCore {
	return entity.AnimeCore{
		Title:    data.Title,
		Synopsis: data.Synopsis,
		Image:    data.Image,
		Genre:    ListGenreRequestToCoreGenre(data.Genre),
	}
}

func GenreRequestToCoreGenre(data GenreRequest) entity.GenreCore {
	return entity.GenreCore{
		Genre:   data.Genre,
		AnimeId: data.AnimeId,
	}
}

func ListGenreRequestToCoreGenre(data []GenreRequest) []entity.GenreCore {
	list := []entity.GenreCore{}
	for _, value := range data {
		result := GenreRequestToCoreGenre(value)
		list = append(list, result)
	}
	return list
}

func GenresRequestToCoreGenre(request GenresRequest) []entity.GenreCore {
    genres := []entity.GenreCore{}
    for _, genre := range request.Genres {
        genres = append(genres, entity.GenreCore{
            Genre:   genre,
            AnimeId: request.AnimeId,
        })
    }
    return genres
}
