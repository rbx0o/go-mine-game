package transfer

import (
	"math"
	"net/http"
	"time"

	"github.com/rbx0o/go-mine-game/domain"
	"github.com/rbx0o/go-mine-game/service"
)

// Здесь будут описаны обработчики с методом GET

/*
pattern:	/game/state
method:		GET
info:		-

succeed:
  - status code:	200 OK
  - response body:	JSON game status + time

failed:
  - status code:	500 InternalServerError
  - response body: 	JSON with error + time
*/
func (h *HTTPHandlers) GetState(response http.ResponseWriter, request *http.Request) {
	state := h.gameService.GetState()
	dto := SuccessResponseDTO[service.GameState]{
		Data:    &state,
		Message: "",
		Time:    time.Now(),
	}
	SendJSON(response, dto, http.StatusOK)
}

/*
pattern:	/enterprise/info
method:		GET
info:		-

succeed:
  - status code:	200 OK
  - response body:	JSON enterprise info + time

failed:
  - status code:	409 Conflict, 500 InternalServerError
  - response body: 	JSON with error + time
*/
func (h *HTTPHandlers) GetIntermediateEnterpriseInfo(response http.ResponseWriter, request *http.Request) {
	err, intermediateInfo := h.gameService.GetIntermediateInfo()

	switch err {
	case service.GameNotRunningYet, service.GameAlreadyFinished:
		dto := ErrorResponseDTO{
			Error: err.Error(),
			Time:  time.Now(),
		}
		SendJSON(response, dto, http.StatusConflict)
		return
	case nil:
		dto := SuccessResponseDTO[service.IntermediateInfo]{
			Data:    intermediateInfo,
			Message: "",
			Time:    time.Now(),
		}
		SendJSON(response, dto, http.StatusOK)
		return
	default:
		dto := ErrorResponseDTO{
			Error: err.Error(),
			Time:  time.Now(),
		}
		SendJSON(response, dto, http.StatusInternalServerError)
		return
	}
}

/*
pattern:	/miners/info
method:		GET
info:		-

succeed:
  - status code:	200 OK
  - response body:	JSON miners info + time

failed:
  - status code:	500 InternalServerError
  - response body: 	JSON with error + time
*/
func (h *HTTPHandlers) GetMinersInfo(response http.ResponseWriter, request *http.Request) {
	minersInfo := h.gameService.GetMinerTypesInfo()

	result := make(map[domain.MinerType]MinerConfigDTO, len(minersInfo))
	for key, value := range minersInfo {
		result[key] = MinerConfigDTO{
			Salary:    value.Salary,
			Energy:    value.Energy,
			CoalCount: value.CoalCount,
			Timeout:   int(math.Round(value.Timeout.Seconds())),
			Progress:  value.Progress,
		}
	}

	dto := SuccessResponseDTO[map[domain.MinerType]MinerConfigDTO]{
		Data:    &result,
		Message: "",
		Time:    time.Now(),
	}
	SendJSON(response, dto, http.StatusOK)
}

/*
pattern:	/equipment
method:		GET
info:		-

succeed:
  - status code:	200 OK
  - response body:	JSON equipment + time

failed:
  - status code:	409 Conflict, 500 InternalServerError
  - response body: 	JSON with error + time
*/
func (h *HTTPHandlers) GetEquipment(response http.ResponseWriter, request *http.Request) {
	err, result := h.gameService.GetEquipmentInfo()

	switch err {
	case service.GameNotRunningYet, service.GameAlreadyFinished:
		dto := ErrorResponseDTO{
			Error: err.Error(),
			Time:  time.Now(),
		}
		SendJSON(response, dto, http.StatusConflict)
		return
	case nil:
		dto := SuccessResponseDTO[map[domain.EquipmentType]bool]{
			Data:    &result,
			Message: "",
			Time:    time.Now(),
		}
		SendJSON(response, dto, http.StatusOK)
		return
	default:
		dto := ErrorResponseDTO{
			Error: err.Error(),
			Time:  time.Now(),
		}
		SendJSON(response, dto, http.StatusInternalServerError)
		return
	}
}

/*
pattern:	/equipment/info
method:		GET
info:		-

succeed:
  - status code:	200 OK
  - response body:	JSON equipment info + time

failed:
  - status code:	500 InternalServerError
  - response body: 	JSON with error + time
*/
func (h *HTTPHandlers) GetEquipmentInfo(response http.ResponseWriter, request *http.Request) {
	result := h.gameService.GetEquipmentTypesInfo()

	dto := SuccessResponseDTO[map[domain.EquipmentType]domain.EquipmentInfo]{
		Data:    &result,
		Message: "",
		Time:    time.Now(),
	}
	SendJSON(response, dto, http.StatusOK)
}
