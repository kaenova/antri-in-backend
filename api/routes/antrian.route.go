package routes

import (
	"antri-in-backend/api/controller"

	"github.com/labstack/echo/v4"
)

func Antrian(e *echo.Echo) *echo.Echo {
	g := e.Group("/antrian")
	g.POST("", controller.AntrianPost)
	return e
}
