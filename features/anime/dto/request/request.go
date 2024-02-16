package request

type AnimeRequest struct {
	Title    string         `json:"title" form:"title"`
	Synopsis string         `json:"synopsis" form:"synopsis"`
	Image    string         `json:"image" form:"image"`
	Genre    []GenreRequest `json:"genre" form:"genre"`
}

type GenreRequest struct {
	Genre   string `json:"genre" form:"genre"`
	AnimeId string `json:"anime_id" form:"anime_id"`
}

type GenresRequest struct {
	Genres  []string `json:"genres" form:"genres"`
	AnimeId string   `json:"anime_id" form:"anime_id"`
}

