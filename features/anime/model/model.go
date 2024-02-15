package model

import (
	"time"

	"gorm.io/gorm"
)

type Anime struct {
	Id        string `gorm:"primaryKey"`
	Title     string `gorm:"not null;unique"`
	Synopsis  string `gorm:"not null"`
	Image     string
	Genre     []Genre `gorm:"foreignKey:GenreId;constraint:OnDelete:CASCADE"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type Genre struct {
	Id        string `gorm:"primaryKey"`
	Genre     string `gorm:"not null"`
	GenreId   string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
