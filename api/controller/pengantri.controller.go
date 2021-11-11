package controller

import (
	"antri-in-backend/api/model"
	"antri-in-backend/entity"
	"net/http"
	"strings"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

func PengantriPost(c echo.Context) error {
	var (
		res entity.Response = entity.CreateResponse()
		req entity.Pengantri
	)

	// Form Input
	namaInput := c.FormValue("nama")
	noHPInput := c.FormValue("no_telp")
	idAntrianInput := c.FormValue("antrian_id")

	// Nomor HP Mau diperiksa?

	// Chcek Input
	if strings.TrimSpace(namaInput) == "" || strings.TrimSpace(noHPInput) == "" {
		res.Message = "Input form tidak valid"
		return c.JSON(res.Status, res)
	}
	antrianID, err := uuid.Parse(idAntrianInput)
	if err != nil {
		res.Message = "Input ID antrian tidak valid"
		return c.JSON(res.Status, res)
	}

	// Binding Data
	req.Nama = namaInput
	req.NoTelp = noHPInput
	req.AntrianID = antrianID

	// Input to DB
	err = model.AddPengantriToAntrian(&req)
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
