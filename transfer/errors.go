package transfer

import (
	"encoding/json"
	"net/http"
	"time"
)

type ErrorDTO struct {
	Message string    `json:"message"`
	Time    time.Time `json:"time"`
}

func (e ErrorDTO) ToString() string {
	b, err := json.MarshalIndent(e, "", "	")
	if err != nil {
		panic(err)
	}
	return string(b)
}

func ErrorToJson(err error, time time.Time, response http.ResponseWriter, status int) {
	errDTO := ErrorDTO{
		Message: err.Error(),
		Time:    time,
	}

	http.Error(response, errDTO.ToString(), status)
}
