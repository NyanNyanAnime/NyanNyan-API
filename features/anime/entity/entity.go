package entity

import (
	"time"

	"gorm.io/gorm"
)

type AnimeCore struct {
	Id        string
	Title     string
	Synopsis  string
	Image     string
	Genre     []GenreCore
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}

type GenreCore struct {
	Id        string
	Genre     string
	GenreId   string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}
