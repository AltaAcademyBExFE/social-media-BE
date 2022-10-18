package main

import (
	"sosmed/config"
	cd "sosmed/features/comment/delivery"
	cr "sosmed/features/comment/repository"
	cs "sosmed/features/comment/services"
	pd "sosmed/features/post/delivery"
	pr "sosmed/features/post/repository"
	ps "sosmed/features/post/services"
	"sosmed/utils/database"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	cfg := config.NewConfig()
	db := database.InitDB(cfg)
	database.MigrateDB(db)

	pRepo := pr.New(db)
	pService := ps.New(pRepo)
	pd.New(e, pService)

	cRepo := cr.New(db)
	cService := cs.New(cRepo)
	cd.New(e, cService)

	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.CORS())
	e.Use(middleware.Logger())

	e.Logger.Fatal(e.Start(":8000"))
}
