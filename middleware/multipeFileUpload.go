package middleware

import (
	"context"
	"fmt"
	dto "moonlay/dto/result"
	"os"
	"path/filepath"

	"net/http"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/labstack/echo/v4"
)

type ImageResult struct {
	PublicID  string
	SecureURL string
}

func UploadMultipleFile(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {

		const MAX_UPLOAD_SIZE = 10 << 20

		if err := c.Request().ParseMultipartForm(MAX_UPLOAD_SIZE); err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}
		if c.Request().ContentLength > MAX_UPLOAD_SIZE {
			response := dto.ErrorResult{Code: http.StatusBadRequest, Message: "Max size in 10mb"}
			return c.JSON(http.StatusBadRequest, response)
		}
		files := c.Request().MultipartForm.File["post_image"]
		if len(files) == 0 {
			c.Set("Error", true)
			return next(c)
		}
		var ctxcld = context.Background()
		var CLOUD_NAME = os.Getenv("CLOUD_NAME")
		var API_KEY = os.Getenv("API_KEY")
		var API_SECRET = os.Getenv("API_SECRET")

		var imageResult []ImageResult
		for _, fileHeader := range files {

			allowedExtensions := map[string]bool{".pdf": true, ".txt": true}
			ext := filepath.Ext(fileHeader.Filename)
			if !allowedExtensions[ext] {
				errorMessage := fmt.Sprintf("File with extension %s is not allowed. Only .pdf and .txt files are allowed.", ext)
				response := dto.ErrorResult{Code: http.StatusBadRequest, Message: errorMessage}
				return c.JSON(http.StatusBadRequest, response)

			}

			var each = ImageResult{}
			file, _ := fileHeader.Open()
			defer file.Close()
			cld, _ := cloudinary.NewFromParams(CLOUD_NAME, API_KEY, API_SECRET)
			resp, err := cld.Upload.Upload(ctxcld, file, uploader.UploadParams{Folder: "MoonLay_Images"})
			if err != nil {
				fmt.Println(err.Error())
			}
			each.PublicID = resp.PublicID
			each.SecureURL = resp.SecureURL
			imageResult = append(imageResult, each)
		}
		// add filename to ctx
		c.Set("dataFiles", imageResult)
		return next(c)
	}
}
