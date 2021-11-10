package routes

import (
	"antri-in-backend/api/controller"
	"antri-in-backend/utils"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Admin(e *echo.Echo) *echo.Echo {
	g := e.Group("/admin")
	g.POST("", controller.AdminPost)

	s := e.Group("/admin")
	s.Use(middleware.JWTWithConfig(utils.JWTconfigAdmin))
	s.GET("", func(c echo.Context) error { return c.String(http.StatusOK, "Bruh") })
	return e
}
