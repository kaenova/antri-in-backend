package api

import (
	"antri-in-backend/api/routes"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/rs/zerolog/log"
)

func Init() *echo.Echo {
	log.Info().Msg("menginisialisasikan server")

	e := echo.New()
	e.Use(middleware.CORS())
	e.Use(middleware.Gzip())
	e.Use(middleware.Logger())
	e = routes.Init(e)

	log.Info().Msg("server terinisialisasi")

	return e
}
