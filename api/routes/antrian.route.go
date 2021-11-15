package routes

import (
	"antri-in-backend/api/controller"
	"antri-in-backend/utils"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Antrian(e *echo.Echo) *echo.Echo {
	// Public
	e.GET("/antrian", controller.AntrianGet)

	// Admin Only
	g := e.Group("/admin/antrian")
	g.Use(middleware.JWTWithConfig(utils.JWTconfigAdmin))
	g.POST("", controller.AntrianPost)
	g.PUT("", controller.AntrianPut)
	g.DELETE("", controller.AntrianDelete)
	g.GET("", controller.AntrianGetAdmin)

	e.POST("/tambah", controller.AntrianTambah, middleware.JWTWithConfig(utils.JWTconfigAdmin))

	return e
}
