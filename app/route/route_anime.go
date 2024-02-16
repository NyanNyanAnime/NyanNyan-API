package route

import (
	"nyannyan/features/anime/controller"
	"nyannyan/features/anime/repository"
	"nyannyan/features/anime/service"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func RouteAnime(e *echo.Group, db *gorm.DB) {
	animeRepository := repository.NewAnimeRepository(db)
	animeService := service.NewAnimeService(animeRepository)
	animeController := controller.NewAnimeController(animeService)

	anime := e.Group("/anime")
	anime.POST("", animeController.CreateAnime)
	anime.GET("", animeController.GetAllAnime)
	anime.GET("/:id", animeController.GetAnimeById)
	anime.PUT("/:id", animeController.UpdateAnimeById)
	anime.DELETE("/:id", animeController.DeleteAnimeById)

	genre := e.Group("/genre")
	genre.POST("", animeController.CreateGenre)
	genre.PATCH("/:id", animeController.UpdateGenreById)
}
