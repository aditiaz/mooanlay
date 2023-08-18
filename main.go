package main

import (
	"fmt"
	"log"
	"moonlay/database"
	"moonlay/pkg/postgresql"
	"moonlay/routes"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	e := echo.New()

	postgresql.DatabaseInit()
	database.RunMigration()
	routes.RouteInit(e.Group("/api/v1"))
	fmt.Println("server running localhost:5000")
	e.Logger.Fatal(e.Start("localhost:5000"))
}
