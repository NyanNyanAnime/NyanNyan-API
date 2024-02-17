package database

import (
	"fmt"

	"nyannyan/features/anime/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"nyannyan/app/config"
)

func InitDBPostgres(cfg *config.AppConfig) *gorm.DB {

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Jakarta",
		cfg.DBHOST, cfg.DBUSER, cfg.DBPASS, cfg.DBNAME, cfg.DBPORT)
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  dsn,
		PreferSimpleProtocol: true,
	}), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	return db
}

func InitMigrationPostgre(db *gorm.DB) {
	db.AutoMigrate(&model.Anime{}, &model.Genre{})
}
