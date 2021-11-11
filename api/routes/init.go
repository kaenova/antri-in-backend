package routes

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
)

func Init(e *echo.Echo) *echo.Echo {

	e = Pengantri(e)
	e = Antrian(e)
	e = Admin(e)
	e = Login(e)

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello world")
	})

	log.Info().Msg("routes terinisialisasi")

	return e
}
