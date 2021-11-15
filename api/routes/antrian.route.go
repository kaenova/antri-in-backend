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
	e.POST("/antrian", controller.AntrianPost, middleware.JWTWithConfig(utils.JWTconfigAdmin))
	e.PUT("/antrian", controller.AntrianPut, middleware.JWTWithConfig(utils.JWTconfigAdmin))
	e.DELETE("/antrian", controller.AntrianDelete, middleware.JWTWithConfig(utils.JWTconfigAdmin))

	e.POST("/tambah", controller.AntrianTambah, middleware.JWTWithConfig(utils.JWTconfigAdmin))

	return e
}
