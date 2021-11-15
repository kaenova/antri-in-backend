package routes

import (
	"antri-in-backend/api/controller"
	"antri-in-backend/utils"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Pengantri(e *echo.Echo) *echo.Echo {
	g := e.Group("/pengantri")
	g.POST("", controller.PengantriPost)
	g.GET("", controller.PengantriGet)

	e.GET("/trace", controller.PengantriTrace, middleware.JWTWithConfig(utils.JWTconfigPengantri))
	return e
}
