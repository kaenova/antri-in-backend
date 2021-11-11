package controller

import (
	"antri-in-backend/api/model"
	"antri-in-backend/entity"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

func AntrianPost(c echo.Context) error {
	var (
		res entity.Response = entity.CreateResponse()
		req entity.Antrian
	)

	// Form Input
	namaInput := c.FormValue("nama")
	deskripsiInput := c.FormValue("deskripsi")

	// Apakah mau di cek role yang buat?

	// Chcek Input
	if strings.TrimSpace(namaInput) == "" || strings.TrimSpace(deskripsiInput) == "" {
		res.Message = "Input form tidak valid"
		return c.JSON(res.Status, res)
	}

	// Binding Data
	req.Nama = namaInput
	req.Deskripsi = deskripsiInput

	// Input to DB
	err := model.AntrianCreate(&req)
	if err != nil {
		res.Status = http.StatusInternalServerError
		res.Message = err.Error()
		return c.JSON(res.Status, res)
	}

	res.Data = req
	res.Message = "Success"
	res.Status = http.StatusOK
	return c.JSON(res.Status, res)
}
