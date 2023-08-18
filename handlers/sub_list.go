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

type handlerSubList struct {
	SubListRepository repositories.SubListRepository
}

func HandlerSubList(sublistRepository repositories.SubListRepository) *handlerSubList {
	return &handlerSubList{sublistRepository}
}

func (h *handlerSubList) FindSubLists(c echo.Context) error {
	sublists, err := h.SubListRepository.GetAllSubLists()
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	postImageSubs, err := h.SubListRepository.GetAllPostImageSubs()
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	postImageSubsMap := make(map[int][]models.PostImageSub)
	for _, postImageSub := range postImageSubs {
		postImageSubsMap[postImageSub.ListID] = append(postImageSubsMap[postImageSub.ListID], postImageSub)
	}

	for i, sublist := range sublists {
		sublists[i].PostImageSub = postImageSubsMap[sublist.ID]
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: sublists})
}

func (h *handlerSubList) GetSubList(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	var sublist models.SubList
	sublist, err := h.SubListRepository.GetSubList(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: sublist})
}
func (h *handlerSubList) DeleteSubList(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	err := h.SubListRepository.DeleteSubList(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: "List deleted successfully"})
}

func (h *handlerSubList) AddSubList(c echo.Context) error {

	request := listdto.SubListRequest{
		Title:     c.FormValue("title"),
		Deskripsi: c.FormValue("deskripsi"),
		ListId:    c.FormValue("list_id"),
	}

	validation := validator.New()
	err := validation.Struct(request)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
	}
	var postSubImage []models.PostImageSub
	dataContex := c.Get("dataFiles")
	filename := dataContex.([]middleware.ImageResult)
	postSubImage = make([]models.PostImageSub, len(filename))
	for i, value := range filename {
		postSubImage[i] = models.PostImageSub{ID: value.PublicID, Image: value.SecureURL}
	}
	id, _ := strconv.Atoi(request.ListId)
	sublist := models.SubList{
		Title:        request.Title,
		Deskripsi:    request.Deskripsi,
		PostImageSub: postSubImage,
		ListId:       id,
		SubListId:    id,
	}

	sublist, err = h.SubListRepository.CreateSubList(sublist)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()})
	}
	return c.JSON(http.StatusOK, dto.SuccessResult{
		Code: http.StatusOK, Data: sublist})
}

func (h *handlerSubList) UpdateSubList(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	request := models.SubList{
		Title:     c.FormValue("title"),
		Deskripsi: c.FormValue("deskripsi"),
	}
	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	sublist, err := h.SubListRepository.GetSubList(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	if request.Title != "" {
		sublist.Title = request.Title
	}
	if request.Deskripsi != "" {
		sublist.Deskripsi = request.Deskripsi
	}

	if len(c.Request().MultipartForm.File["post_image"]) > 0 {

		var postImage []models.PostImageSub
		dataContex := c.Get("dataFiles")
		filename := dataContex.([]middleware.ImageResult)
		postImage = make([]models.PostImageSub, len(filename))
		for i, value := range filename {
			postImage[i] = models.PostImageSub{ID: value.PublicID, Image: value.SecureURL}
		}

		sublist.PostImageSub = postImage
	}

	data, err := h.SubListRepository.UpdateSubList(sublist)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()})
	}
	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: data})
}
