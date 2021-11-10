package routes

import (
	"antri-in-backend/api/controller"

	"github.com/labstack/echo/v4"
)

func Login(e *echo.Echo) *echo.Echo {
	g := e.Group("/login")
	g.POST("", controller.LoginPost)
	return e
}
