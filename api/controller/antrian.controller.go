package controller

import (
	"antri-in-backend/api/model"
	"antri-in-backend/entity"
	"antri-in-backend/utils"
	"net/http"
	"strconv"
	"strings"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
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
	userData := c.Get("user").(*jwt.Token)
	claims := userData.Claims.(*utils.JWTCustomClaimsAdmin)
	roles := claims.Role
	if roles != entity.ROLES_SUPER_USER {
		res.Message = "Not Authorized"
		return c.JSON(res.Status, res)
	}

	// Chcek Input
	if strings.TrimSpace(namaInput) == "" || strings.TrimSpace(deskripsiInput) == "" {
		res.Message = "Input form tidak valid"
		return c.JSON(res.Status, res)
	}

	// Binding Data
	req.Nama = namaInput
	req.Deskripsi = deskripsiInput

	// Cek apakah ada nama yang sama?
	used, err := model.AntrianSearchNameExactUsed(req.Nama)
	if err != nil {
		res.Status = http.StatusInternalServerError
		res.Message = "Internal Server Error"
		return c.JSON(res.Status, res)
	}

	if used {
		res.Message = "Antrian sudah digunakan"
		return c.JSON(res.Status, res)
	}

	// Input to DB
	err = model.AntrianCreate(&req)
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

func AntrianGet(c echo.Context) error {
	var (
		res  entity.Response = entity.CreateResponse()
		done bool            = false
		err  error
	)

	/*
		Priority
		1. Search by ID
		2. Search by name
		3. Get ALl
	*/

	// Search by ID
	idInput := c.QueryParam("id")
	if strings.TrimSpace(idInput) != "" && !done {
		idAntrian, err := uuid.Parse(idInput)
		if err != nil {
			res.Message = "ID Tidak Valid"
			return c.JSON(res.Status, res)
		}
		res.Data, err = model.AntrianbyID(idAntrian)
		if err != nil {
			res.Status = http.StatusInternalServerError
			res.Message = err.Error()
			return c.JSON(res.Status, res)
		}
		done = true
	}

	// Search by name
	searchInput := c.QueryParam("search")
	if strings.TrimSpace(searchInput) != "" && !done {
		res.Data, err = model.AntrianSearchName(searchInput)
		if err != nil {
			res.Status = http.StatusInternalServerError
			res.Message = err.Error()
			return c.JSON(res.Status, res)
		}
		done = true
	}

	// Get All
	if !done {
		res.Data, err = model.AntrianGetAll()
		if err != nil {
			res.Status = http.StatusInternalServerError
			res.Message = err.Error()
			return c.JSON(res.Status, res)
		}
		done = true
	}

	res.Status = http.StatusOK
	res.Message = "Success"
	return c.JSON(res.Status, res)
}

func AntrianDelete(c echo.Context) error {
	var (
		res entity.Response = entity.CreateResponse()
		err error
	)

	// Cek apakah super user?
	userData := c.Get("user").(*jwt.Token)
	claims := userData.Claims.(*utils.JWTCustomClaimsAdmin)
	roles := claims.Role
	if roles != entity.ROLES_SUPER_USER {
		res.Message = "Not Authorized"
		return c.JSON(res.Status, res)
	}

	idInput := c.QueryParam("id")
	idAntrian, err := uuid.Parse(idInput)
	if err != nil {
		res.Message = "ID not valid"
		return c.JSON(res.Status, res)
	}

	res.Data, err = model.AntrianDelete(idAntrian)
	if err != nil {
		res.Status = http.StatusInternalServerError
		res.Message = err.Error()
		return c.JSON(res.Status, res)
	}

	res.Status = http.StatusOK
	res.Message = "Success"
	return c.JSON(res.Status, res)
}

func AntrianPut(c echo.Context) error {
	var (
		res entity.Response = entity.CreateResponse()
		req entity.Antrian
	)

	// Cek apakah super user?
	userData := c.Get("user").(*jwt.Token)
	claims := userData.Claims.(*utils.JWTCustomClaimsAdmin)
	roles := claims.Role
	if roles != entity.ROLES_SUPER_USER {
		res.Message = "Not Authorized"
		return c.JSON(res.Status, res)
	}

	idInput := c.QueryParam("id")
	idAntrian, err := uuid.Parse(idInput)
	if err != nil {
		res.Message = "ID not valid"
		return c.JSON(res.Status, res)
	}

	// Form Input
	namaInput := c.FormValue("nama")
	deskripsiInput := c.FormValue("deskripsi")

	// Chcek Input
	if strings.TrimSpace(namaInput) == "" || strings.TrimSpace(deskripsiInput) == "" {
		res.Message = "Input form tidak valid"
		return c.JSON(res.Status, res)
	}

	// Binding Data
	req.Nama = namaInput
	req.Deskripsi = deskripsiInput
	req.ID = idAntrian

	// Input to DB
	err = model.AntrianUbah(&req)
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

func AntrianTambah(c echo.Context) error {
	var (
		res entity.Response = entity.CreateResponse()
		req entity.Antrian
	)

	idInput := c.QueryParam("id")
	idAntrian, err := uuid.Parse(idInput)
	if err != nil {
		res.Message = "ID not valid"
		return c.JSON(res.Status, res)
	}

	// Binding
	req.ID = idAntrian

	// Input to DB
	err = model.TambahNomorAntrian(&req)
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

func AntrianGetAdmin(c echo.Context) error {
	var (
		res  entity.Response = entity.CreateResponse()
		done bool            = false
		err  error
	)

	// Get Pengantri
	antrianInput := c.QueryParam("pengantri")
	idInput := c.QueryParam("id")
	if strings.TrimSpace(antrianInput) != "" && strings.TrimSpace(idInput) != "" && !done {
		_, err1 := strconv.ParseBool(antrianInput)
		idAntrian, err2 := uuid.Parse(idInput)
		if err1 != nil || err2 != nil {
			res.Message = "Kesalahan parameter, harap cek dokumentasi API"
			return c.JSON(res.Status, res)
		}
		res.Data, err = model.AmbilPengantribyAntrianID(idAntrian)
		if err != nil {
			res.Message = err.Error()
			res.Status = http.StatusInternalServerError
			return c.JSON(res.Status, res)
		}
		done = true
	}

	res.Message = "Success"
	res.Status = http.StatusOK
	return c.JSON(res.Status, res)
}
