package routes

import (
	"moonlay/handlers"
	"moonlay/middleware"
	"moonlay/pkg/postgresql"
	"moonlay/repositories"

	"github.com/labstack/echo/v4"
)

func ListRoutes(e *echo.Group) {
	repo := repositories.MakeRepository(postgresql.DB)
	h := handlers.HandlerList(repo)

	e.POST("/list", middleware.UploadMultipleFile(h.AddList))
	e.GET("/list/:id", h.GetList)
	e.GET("/lists", h.FindLists)
	e.DELETE("/list/:id", h.DeleteList)
	e.PATCH("/updatelist/:id", middleware.UploadMultipleFile(h.UpdateList))
	e.GET("/searchlist", h.SearchLists)

}
