package controller

import (
	"antri-in-backend/api/model"
	"antri-in-backend/entity"
	"antri-in-backend/utils"
	"net/http"
	"net/mail"

	"github.com/labstack/echo/v4"
)

func LoginPost(c echo.Context) error {
	var (
		res entity.Response = entity.CreateResponse()
	)

	// Get Request
	emailForm := c.FormValue("email")
	password := c.FormValue("password")

	// Email Check
	if _, err := mail.ParseAddress(emailForm); err != nil {
		res.Message = "Email tidak valid"
		return c.JSON(res.Status, res)
	}

	// Check Email
	used, obj := model.AdminEmailIsUsed(emailForm)
	if !used {
		res.Message = "Email tidak terdaftar"
		return c.JSON(res.Status, res)
	}

	// Password Validation
	valid := utils.CompareHashPassword(password, obj.Password)
	if !valid {
		res.Message = "Password salah"
		return c.JSON(res.Status, res)
	}

	// Creating JWT
	token, err := utils.GenerateTokenAdmin(obj.Nama, obj.Role, obj.Email, obj.ID.String())
	if err != nil {
		res.Status = http.StatusInternalServerError
		res.Message = "Error when generating JWT"
		return c.JSON(res.Status, res)
	}

	res.Data = map[string]string{"token": token, "nama": obj.Nama, "email": obj.Email, "role": obj.Role}
	res.Status = http.StatusOK
	res.Message = "Success"
	return c.JSON(res.Status, res)
}
