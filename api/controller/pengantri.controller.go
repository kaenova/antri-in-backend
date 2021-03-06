package controller

import (
	"antri-in-backend/api/model"
	"antri-in-backend/entity"
	"antri-in-backend/utils"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt"
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

	// Chcek Input
	if strings.TrimSpace(namaInput) == "" || strings.TrimSpace(noHPInput) == "" {
		res.Message = "Input form tidak valid"
		return c.JSON(res.Status, res)
	}

	// Periksa No HP
	if _, used := model.NoHPIsUsed(noHPInput); used {
		res.Message = "No Hp sudah digunakan"
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

	token, err := utils.GenerateTokenPengantri(req.ID.String(), req.Nama, req.AntrianID.String(), req.NoAntrian)
	if err != nil {
		res.Status = http.StatusInternalServerError
		res.Message = err.Error()
		return c.JSON(res.Status, res)
	}

	res.Data = map[string]string{"token": token}
	res.Message = "Success"
	res.Status = http.StatusOK
	return c.JSON(res.Status, res)
}

func PengantriGet(c echo.Context) error {
	var (
		res entity.Response = entity.CreateResponse()
	)

	// Query Param Input
	noHPInput := c.QueryParam("telp")

	// Get Token by NO HP
	if strings.TrimSpace(noHPInput) == "" {
		res.Message = "No HP Tidak Valid"
		return c.JSON(res.Status, res)
	}

	data, used := model.NoHPIsUsed(noHPInput)
	if !used {
		res.Message = "No HP Tidak Terdaftar"
		return c.JSON(res.Status, res)
	}

	token, err := utils.GenerateTokenPengantri(data.ID.String(), data.Nama, data.AntrianID.String(), data.NoAntrian)
	if err != nil {
		res.Status = http.StatusInternalServerError
		res.Message = err.Error()
		return c.JSON(res.Status, res)
	}

	res.Data = map[string]string{"token": token}
	res.Message = "Success"
	res.Status = http.StatusOK
	return c.JSON(res.Status, res)
}

func PengantriTrace(c echo.Context) error {
	var (
		res entity.Response = entity.CreateResponse()
		err error
	)

	// Cek apakah super user?
	userData := c.Get("user").(*jwt.Token)
	claims := userData.Claims.(*utils.JWTCustomClaimsPengantri)
	idAntriString := claims.IdAntrian
	idAntrian, err := uuid.Parse(idAntriString)
	if err != nil {
		res.Message = "Tidak Valid"
		return c.JSON(res.Status, res)
	}

	res.Data, err = model.AntrianbyID(idAntrian)
	if err != nil {
		res.Status = http.StatusInternalServerError
		res.Message = err.Error()
		return c.JSON(res.Status, res)
	}

	res.Data = echo.Map{"antrian": res.Data, "no_antrian_pengantri": claims.NoAntrian}
	res.Status = http.StatusOK
	res.Message = "Success"
	return c.JSON(res.Status, res)
}

func PengantriDelete(c echo.Context) error {
	var (
		res             entity.Response = entity.CreateResponse()
		err, err1, err2 error
		idString        string
		idUser          uuid.UUID
		jwtData         *jwt.Token
	)
	// Query Input Token Or Ppengantri ID
	tokenOrIDInput := c.FormValue("token_id")
	if strings.TrimSpace(tokenOrIDInput) == "" {
		res.Message = "Invalid Input"
		return c.JSON(res.Status, res)
	}

	// Menggunakan Token
	jwtData, err1 = utils.JWTTokenFromStringPengantri(tokenOrIDInput)
	idUser, err2 = uuid.Parse(tokenOrIDInput)
	if err1 != nil && err2 != nil {
		res.Message = err1.Error()
		return c.JSON(res.Status, res)
	}

	if err2 != nil {
		claims := jwtData.Claims.(*utils.JWTCustomClaimsPengantri)
		// DONT TOUCH IT
		idString = claims.Id
		idUser, err = uuid.Parse(idString)
		if err != nil {
			res.Message = "Invalid Input"
			return c.JSON(res.Status, res)
		}
	}

	err = model.DeletePengantri(idUser)
	if err != nil {
		res.Status = http.StatusInternalServerError
		res.Message = err.Error()
		return c.JSON(res.Status, res)
	}

	res.Message = "Success"
	res.Status = http.StatusOK
	return c.JSON(res.Status, res)
}
