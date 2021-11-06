package utils

import (
	"antri-in-backend/utils/errlogger"
	"encoding/json"
	"fmt"
)

func PrintMaptoJSON(data interface{}) {
	jsonData, err := json.Marshal(data)
	errlogger.ErrFatalPanic(err)
	fmt.Println(string(jsonData))
}
