package controller

import (
	"antri-in-backend/entity"

	"github.com/labstack/echo/v4"
)

func AntrianPost(c echo.Context) error {
	var (
		res entity.Response = entity.CreateResponse()
	)

	return c.JSON(res.Status, res)
}
