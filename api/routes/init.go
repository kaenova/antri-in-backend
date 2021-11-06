package routes

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
)

func Init(e *echo.Echo) *echo.Echo {

	// e = Auth(e)
	// e = Produk(e)
	// e = Transaction(e)

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello world")
	})

	log.Info().Msg("routes terinisialisasi")

	return e
}
