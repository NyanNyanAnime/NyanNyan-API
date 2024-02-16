package model

import (
	"time"

	"gorm.io/gorm"
)

type Anime struct {
	Id        string `gorm:"primaryKey;type:varchar(191)"`
	Title     string `gorm:"not null;unique"`
	Synopsis  string `gorm:"not null"`
	Image     string
	Genre     []Genre `gorm:"foreignKey:AnimeId;constraint:OnDelete:CASCADE"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type Genre struct {
	Id        string `gorm:"primaryKey;type:varchar(191)"`
	Genre     string `gorm:"not null"`
	AnimeId   string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
