package main

import (
	"sosmed/config"
	ud "sosmed/features/user/delivery"
	ur "sosmed/features/user/repository"
	us "sosmed/features/user/services"
	"sosmed/utils/database"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	cfg := config.NewConfig()
	db := database.InitDB(cfg)

	uRepo := ur.New(db)
	uService := us.New(uRepo)
	ud.New(e, uService)

	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.CORS())
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))

	e.Logger.Fatal(e.Start(":8000"))
}
