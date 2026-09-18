package transfer

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type ErrorResponseDTO struct {
	Error string    `json:"error,omitempty"`
	Time  time.Time `json:"time"`
}

type SuccessResponseDTO[T any] struct {
	Data    *T        `json:"data,omitempty"`
	Message string    `json:"message,omitempty"`
	Time    time.Time `json:"time"`
}

/*
SendJSON конвертирует dto в JSON и пытается отправить с указанным статус кодом
*/
func SendJSON(response http.ResponseWriter, dto any, status int) {
	response.Header().Set("Content-Type", "application/json")

	b, err := json.MarshalIndent(dto, "", "\t")
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		return
	}

	response.WriteHeader(status)
	if _, err := response.Write(b); err != nil {
		fmt.Printf("HTTP write header error: %v", err)
		return
	}
}
