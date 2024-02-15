package response

type AnimeResponse struct {
	Id       string          `json:"id"`
	Title    string          `json:"title"`
	Synopsis string          `json:"synopsis"`
	Image    string          `json:"image"`
	Genre    []GenreResponse `json:"genre"`
}

type GenreResponse struct {
	Genre string `json:"genre"`
}
