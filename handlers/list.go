package handlers

import (
	listdto "moonlay/dto/list"
	dto "moonlay/dto/result"
	"moonlay/middleware"
	"moonlay/models"
	"moonlay/repositories"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type handlerList struct {
	ListRepository repositories.ListRepository
}

func HandlerList(listRepository repositories.ListRepository) *handlerList {
	return &handlerList{listRepository}
}

func (h *handlerList) FindLists(c echo.Context) error {
	lists, err := h.ListRepository.GetAllLists()
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}
	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: lists})
}

func (h *handlerList) GetList(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	var list models.List
	list, err := h.ListRepository.GetList(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: list})
}
func (h *handlerList) DeleteList(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	err := h.ListRepository.DeleteList(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: "List deleted successfully"})
}

func (h *handlerList) AddList(c echo.Context) error {

	request := listdto.ListRequest{
		Title:     c.FormValue("title"),
		Deskripsi: c.FormValue("deskripsi"),
	}

	validation := validator.New()
	err := validation.Struct(request)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
	}
	var postImage []models.PostImage
	dataContex := c.Get("dataFiles")
	filename := dataContex.([]middleware.ImageResult)
	postImage = make([]models.PostImage, len(filename))
	for i, value := range filename {
		postImage[i] = models.PostImage{ID: value.PublicID, Image: value.SecureURL, ListID: request.ID}
	}

	list := models.List{
		Title:     request.Title,
		Deskripsi: request.Deskripsi,
		PostImage: postImage,
	}

	list, err = h.ListRepository.CreateList(list)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()})
	}
	return c.JSON(http.StatusOK, dto.SuccessResult{
		Code: http.StatusOK, Data: list})
}

func (h *handlerList) UpdateList(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	request := models.List{
		Title:     c.FormValue("title"),
		Deskripsi: c.FormValue("deskripsi"),
	}
	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	list, err := h.ListRepository.GetList(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	if request.Title != "" {
		list.Title = request.Title
	}
	if request.Deskripsi != "" {
		list.Deskripsi = request.Deskripsi
	}

	if len(c.Request().MultipartForm.File["post_image"]) > 0 {

		var postImage []models.PostImage
		dataContex := c.Get("dataFiles")
		filename := dataContex.([]middleware.ImageResult)
		postImage = make([]models.PostImage, len(filename))
		for i, value := range filename {
			postImage[i] = models.PostImage{ID: value.PublicID, Image: value.SecureURL}
		}

		list.PostImage = postImage
	}

	data, err := h.ListRepository.UpdateList(list)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()})
	}
	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: data})
}

func (h *handlerList) SearchLists(c echo.Context) error {
	searchQuery := c.QueryParam("q")

	if searchQuery == "" {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: "Search query is required."})
	}

	lists, err := h.ListRepository.SearchLists(searchQuery)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()})
	}

	response := struct {
		Lists []models.List `json:"lists"`
	}{
		Lists: lists,
	}

	return c.JSON(http.StatusOK, response)
}
