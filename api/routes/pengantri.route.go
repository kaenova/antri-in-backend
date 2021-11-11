package routes

import (
	"antri-in-backend/api/controller"

	"github.com/labstack/echo/v4"
)

func Pengantri(e *echo.Echo) *echo.Echo {
	g := e.Group("/pengantri")
	g.POST("", controller.PengantriPost)
	return e
}
