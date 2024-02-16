package response

type AnimeResponse struct {
	Id       string          `json:"id"`
	Title    string          `json:"title"`
	Synopsis string          `json:"synopsis"`
	Image    string          `json:"image"`
	Genres   []GenreResponse `json:"genres"`
}

type GenreResponse struct {
	Genre string `json:"genre"`
}
