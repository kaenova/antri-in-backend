package routes

import (
	"antri-in-backend/api/controller"
	"antri-in-backend/utils"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Admin(e *echo.Echo) *echo.Echo {
	g := e.Group("/admin")
	g.POST("", controller.AdminPost)

	s := e.Group("/admin/request")
	s.Use(middleware.JWTWithConfig(utils.JWTconfigAdmin))
	s.GET("", controller.AdminRequestGet)
	s.POST("", controller.AdminRequestPost)
	s.DELETE("", controller.AdminRequestDelete)

	return e
}
