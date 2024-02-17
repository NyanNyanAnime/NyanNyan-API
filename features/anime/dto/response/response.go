package response

type AnimeResponse struct {
	Id        string          `json:"id"`
	Title     string          `json:"title"`
	Synopsis  string          `json:"synopsis"`
	Type      string          `json:"type"`
	Episodes  int             `json:"episodes"`
	Premiered string          `json:"premiered"`
	Aired     string          `json:"aired"`
	Studios   string          `json:"studios"`
	Duration  string          `json:"duration"`
	Rating    string          `json:"rating"`
	Image     string          `json:"image"`
	Genres    []GenreResponse `json:"genres"`
}

type GenreResponse struct {
	Genre string `json:"genre"`
}
