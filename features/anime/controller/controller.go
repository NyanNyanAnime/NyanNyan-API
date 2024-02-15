package controller

import (
	"net/http"
	"nyannyan/features/anime/dto/request"
	"nyannyan/features/anime/dto/response"
	"nyannyan/features/anime/entity"
	"nyannyan/utils/constanta"
	"nyannyan/utils/helper"
	"strings"

	"github.com/labstack/echo/v4"
)

type animeController struct {
	animeService entity.AnimeServiceInterface
}

func NewAnimeController(anime entity.AnimeServiceInterface) *animeController {
	return &animeController{
		animeService: anime,
	}
}

func (ac *animeController) CreateAnime(e echo.Context) error {

	input := request.AnimeRequest{}
	err := helper.BindFormData(e, &input)
	if err != nil {
		return e.JSON(http.StatusBadRequest, helper.ErrorResponse(err.Error()))
	}

	image, err := e.FormFile("image")
	if err != nil {
		if err == http.ErrMissingFile {
			return e.JSON(http.StatusBadRequest, helper.ErrorResponse(constanta.ERROR_EMPTY_FILE))
		}
		return e.JSON(http.StatusBadRequest, helper.ErrorResponse("failed upload file"))
	}

	request := request.AnimeRequestToCoreAnime(input)
	err = ac.animeService.CreateAnime(image, request)
	if err != nil {
		return e.JSON(http.StatusBadRequest, helper.ErrorResponse(err.Error()))
	}

	return e.JSON(http.StatusCreated, helper.SuccessResponse(constanta.SUCCESS_CREATE_DATA))
}

func (ah *animeController) GetAllAnime(e echo.Context) error {
	result, err := ah.animeService.GetAllAnime()
	if err != nil {
		return e.JSON(http.StatusInternalServerError, helper.ErrorResponse(err.Error()))
	}

	if len(result) == 0 {
		return e.JSON(http.StatusOK, helper.SuccessResponse(constanta.SUCCESS_NULL))
	}

	response := response.ListCoreAnimeToAnimeResponse(result)
	return e.JSON(http.StatusOK, helper.SuccessWithDataResponse(constanta.SUCCESS_GET_DATA, response))
}

func (ac *animeController) GetAnimeById(e echo.Context) error {

	id := e.Param("id")
	result, err := ac.animeService.GetAnimeById(id)
	if err != nil {
		if strings.Contains(err.Error(), constanta.ERROR_RECORD_NOT_FOUND) {
			return e.JSON(http.StatusNotFound, helper.ErrorResponse(constanta.ERROR_DATA_NOT_FOUND))
		}
		return e.JSON(http.StatusInternalServerError, helper.ErrorResponse(err.Error()))
	}

	response := response.CoreAnimeToAnimeResponse(result)
	return e.JSON(http.StatusOK, helper.SuccessWithDataResponse(constanta.SUCCESS_GET_DATA, response))
}

func (ac *animeController) UpdateAnimeById(e echo.Context) error {
	input := request.AnimeRequest{}

	err := helper.BindFormData(e, &input)
	if err != nil {
		return e.JSON(http.StatusBadRequest, helper.ErrorResponse(err.Error()))
	}

	id := e.Param("id")
	image, _ := e.FormFile("image")

	request := request.AnimeRequestToCoreAnime(input)
	err = ac.animeService.UpdateAnimeById(id, image, request)
	if err != nil {
		if strings.Contains(err.Error(), constanta.ERROR_RECORD_NOT_FOUND) {
			return e.JSON(http.StatusNotFound, helper.ErrorResponse(constanta.ERROR_DATA_NOT_FOUND))
		}
		return e.JSON(http.StatusInternalServerError, helper.ErrorResponse(err.Error()))
	}

	return e.JSON(http.StatusOK, helper.SuccessResponse(constanta.SUCCESS_UPDATE_DATA))
}

func (ac *animeController) DeleteAnimeById(e echo.Context) error {
	id := e.Param("id")
	err := ac.animeService.DeleteAnimeById(id)
	if err != nil {
		if strings.Contains(err.Error(), constanta.ERROR_RECORD_NOT_FOUND) {
			return e.JSON(http.StatusNotFound, helper.ErrorResponse(constanta.ERROR_DATA_NOT_FOUND))
		}
		return e.JSON(http.StatusInternalServerError, helper.ErrorResponse(err.Error()))
	}

	return e.JSON(http.StatusOK, helper.SuccessResponse(constanta.SUCCESS_DELETE_DATA))
}