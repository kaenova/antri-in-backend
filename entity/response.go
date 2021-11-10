package entity

import (
	"net/http"
)

type Response struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func CreateResponse() (r Response) {
	r.Status = http.StatusBadRequest
	r.Message = "Bad Request"
	return
}
