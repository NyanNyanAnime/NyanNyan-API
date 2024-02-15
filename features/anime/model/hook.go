package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func (anime *Anime) BeforeCreate(tx *gorm.DB) (err error) {
	newUuid := uuid.New()
	anime.Id = newUuid.String()

	return nil
}

func (genre *Genre) BeforeCreate(tx *gorm.DB) (err error) {
	newUuid := uuid.New()
	genre.Id = newUuid.String()

	return nil
}
