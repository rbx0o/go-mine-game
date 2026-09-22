package transfer

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/rbx0o/go-mine-game/domain"
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

type MinerTypeDTO struct {
	Type string `json:"type"`
}

type GameResultDTO struct {
	Balance         domain.Coal                    `json:"balance"`
	StartTime       time.Time                      `json:"start_time"`
	EndTime         time.Time                      `json:"end_time"`
	DurationTime    float64                        `json:"duration_time"`
	EndedAuto       bool                           `json:"ended_time"`
	ResultEquipment map[domain.EquipmentType]bool  `json:"result_equipment"`
	ResultMiners    map[domain.ID]domain.MinerInfo `json:"result_miners"`
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

/*
GetJSON получение body из запроса и запись в переданную DTO
*/
func GetFromJSON[T any](request *http.Request, dto *T) error {
	body, err := io.ReadAll(request.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, dto)
	if err != nil {
		return err
	}

	return nil
}
