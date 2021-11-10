package routes

import (
	"antri-in-backend/api/controller"

	"github.com/labstack/echo/v4"
)

func Admin(e *echo.Echo) *echo.Echo {
	g := e.Group("/admin")
	g.POST("", controller.AdminPost)
	return e
}
