package controller

import (
	"antri-in-backend/api/model"
	"antri-in-backend/entity"
	"antri-in-backend/utils"
	"net/http"
	"net/mail"
	"strings"

	"github.com/labstack/echo/v4"
)

func AdminPost(c echo.Context) error {
	var (
		res entity.Response = entity.CreateResponse()
		req entity.Admin
	)

	// Get Request
	namaForm := c.FormValue("nama")
	emailForm := c.FormValue("email")
	password := c.FormValue("password")

	// Name Check
	if strings.TrimSpace(namaForm) == "" {
		res.Message = "Nama tidak boleh kosong"
		return c.JSON(res.Status, res)
	}

	// Email Check
	if _, err := mail.ParseAddress(emailForm); err != nil {
		res.Message = "Email tidak valid"
		return c.JSON(res.Status, res)
	}

	// Password Check

	// Hash password
	passwordHashed, err := utils.HashPassword(password)
	if err != nil {
		res.Message = "Internal Server Error"
		res.Status = http.StatusInternalServerError
		return c.JSON(res.Status, res)
	}

	// Assign to entity
	req.Email = emailForm
	req.Nama = namaForm
	req.Password = passwordHashed
	req.Role = entity.ROLES_ADMIN
	req.IsActive = false

	// Check Email
	used, _ := model.AdminEmailIsUsed(req.Email)
	if used {
		res.Message = "Email sudah digunakan"
		return c.JSON(res.Status, res)
	}

	// Inputting Data
	data, err := model.AddAdmin(req)
	if err != nil {
		res.Status = http.StatusInternalServerError
		res.Message = err.Error()
		return c.JSON(res.Status, res)
	}

	// Hilangkan data yang berbahaya
	data.Password = ""
	data.Role = ""

	res.Data = data
	res.Status = http.StatusOK
	res.Message = "Success"
	return c.JSON(res.Status, res)
}
