package controller

import (
	"net/http"
	"nyannyan/features/anime/dto/request"
	"nyannyan/features/anime/dto/response"
	"nyannyan/features/anime/entity"
	"nyannyan/utils/constanta"
	"nyannyan/utils/helper"
	"strconv"
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

func (ac *animeController) CreateGenre(e echo.Context) error {

	input := request.GenresRequest{}
	err := e.Bind(&input)
	if err != nil {
		return e.JSON(http.StatusBadRequest, helper.ErrorResponse(err.Error()))
	}

	request := request.GenresRequestToCoreGenre(input)
	err = ac.animeService.CreateGenre(request)
	if err != nil {
		return e.JSON(http.StatusBadRequest, helper.ErrorResponse(err.Error()))
	}

	return e.JSON(http.StatusCreated, helper.SuccessResponse(constanta.SUCCESS_CREATE_DATA))
}

func (ac *animeController) CreateAnime(e echo.Context) error {

	input := request.AnimeRequest{}
	err := e.Bind(&input)
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

	search := e.QueryParam("search")
	page, _ := strconv.Atoi(e.QueryParam("page"))
	limit, _ := strconv.Atoi(e.QueryParam("limit"))

	anime, paginationInfo, count, err := ah.animeService.GetAllAnime(page, limit, search)
	if err != nil {
		return e.JSON(http.StatusInternalServerError, helper.ErrorResponse(err.Error()))
	}

	if len(anime) == 0 {
		return e.JSON(http.StatusOK, helper.SuccessResponse(constanta.SUCCESS_NULL))
	}

	response := response.ListCoreAnimeToAnimeResponse(anime)
	return e.JSON(http.StatusOK, helper.SuccessWithPagnationAndCount(constanta.SUCCESS_GET_DATA, response, paginationInfo, count))
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

func (uh *animeController) UpdateGenreById(e echo.Context) error {
	input := request.GenreRequest{}

	errBind := helper.DecodeJSON(e, &input)
	if errBind != nil {
		return e.JSON(http.StatusBadRequest, helper.ErrorResponse(errBind.Error()))
	}

	request := request.GenreRequestToCoreGenre(input)

	id := e.Param("id")
	errUpdate := uh.animeService.UpdateGenreById(id, request)
	if errUpdate != nil {
		return e.JSON(http.StatusBadRequest, helper.ErrorResponse(errUpdate.Error()))
	}

	return e.JSON(http.StatusOK, helper.SuccessResponse(constanta.SUCCESS_UPDATE_DATA))

}

func (ac *animeController) DeleteGenreById(e echo.Context) error {
	id := e.Param("id")
	err := ac.animeService.DeleteGenreById(id)
	if err != nil {
		if strings.Contains(err.Error(), constanta.ERROR_RECORD_NOT_FOUND) {
			return e.JSON(http.StatusNotFound, helper.ErrorResponse(constanta.ERROR_DATA_NOT_FOUND))
		}
		return e.JSON(http.StatusInternalServerError, helper.ErrorResponse(err.Error()))
	}

	return e.JSON(http.StatusOK, helper.SuccessResponse(constanta.SUCCESS_DELETE_DATA))
}