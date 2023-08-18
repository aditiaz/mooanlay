package routes

import (
	"moonlay/handlers"
	"moonlay/middleware"
	"moonlay/pkg/postgresql"
	"moonlay/repositories"

	"github.com/labstack/echo/v4"
)

func SubListRoutes(e *echo.Group) {
	repo := repositories.MakeRepository(postgresql.DB)
	h := handlers.HandlerSubList(repo)

	e.POST("/sublist", middleware.UploadMultipleFile(h.AddSubList))
	e.GET("/sublist/:id", h.GetSubList)
	e.GET("/sublists", h.FindSubLists)
	e.DELETE("/sublist/:id", h.DeleteSubList)
	e.PATCH("/updatesublist/:id", middleware.UploadMultipleFile(h.UpdateSubList))
	e.GET("/searchsublist", h.SearchSubLists)

}
