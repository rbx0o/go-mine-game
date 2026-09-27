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
pattern:	/game/state
method:		GET
info:		-

succeed:
  - status code:	200 OK
  - response body:	JSON game result + time

failed:
  - status code:	409 Conflict, 500 InternalServerError
  - response body: 	JSON with error + time
*/
func (h *HTTPHandlers) GetGameResult(response http.ResponseWriter, request *http.Request) {
	err, result := h.gameService.GetGameResult()

	switch err {
	case service.GameNotRunningYet, service.GameNotEnd:
		dto := ErrorResponseDTO{
			Error: err.Error(),
			Time:  time.Now(),
		}
		SendJSON(response, dto, http.StatusConflict)
		return
	case nil:
		tempDTO := GameResultDTO{
			Balance:         result.Balance,
			StartTime:       result.StartTime,
			EndTime:         result.EndTime,
			DurationTime:    int(math.Round(result.DurationTime.Seconds())),
			EndedAuto:       result.EndedAuto,
			ResultEquipment: result.ResultEquipment,
			ResultMiners:    result.ResultMiners,
		}

		dto := SuccessResponseDTO[GameResultDTO]{
			Data:    &tempDTO,
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
pattern:	/miners/all
method:		GET
info:		-

succeed:
  - status code:	200 OK
  - response body:	JSON miners info + time

failed:
  - status code:	500 InternalServerError
  - response body: 	JSON with error + time
*/
func (h *HTTPHandlers) GetAllMiners(response http.ResponseWriter, request *http.Request) {
	err, result := h.gameService.GetAllMiners()

	switch err {
	case service.GameNotRunningYet, service.GameAlreadyFinished:
		dto := ErrorResponseDTO{
			Error: err.Error(),
			Time:  time.Now(),
		}
		SendJSON(response, dto, http.StatusConflict)
		return
	case nil:
		dto := SuccessResponseDTO[map[domain.ID]domain.MinerInfo]{
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
pattern:	/miners/inactive
method:		GET
info:		-

succeed:
  - status code:	200 Ok
  - response body:	JSON inactive miners + time

failed:
  - status code:	400 Bad Request, 409 Conflict, 500 InternalServerError
  - response body: 	JSON with error + time
*/
func (h *HTTPHandlers) GetInactiveMiners(response http.ResponseWriter, request *http.Request) {
	if request.URL.Query().Has("type") {
		minerType := request.URL.Query().Get("type")

		if minerType == "" {
			errorDTO := ErrorResponseDTO{
				Error: "The query must contain valid parameter",
				Time:  time.Now(),
			}
			SendJSON(response, errorDTO, http.StatusBadRequest)
			return
		} else {
			err, inactiveMiners := h.gameService.GetInactiveMinersFilter(domain.MinerType(minerType))

			switch err {
			case service.GameNotRunningYet, service.GameAlreadyFinished:
				errorDTO := ErrorResponseDTO{
					Error: err.Error(),
					Time:  time.Now(),
				}
				SendJSON(response, errorDTO, http.StatusConflict)
				return
			case service.MinerTypeNotFound:
				errorDTO := ErrorResponseDTO{
					Error: err.Error(),
					Time:  time.Now(),
				}
				SendJSON(response, errorDTO, http.StatusBadRequest)
				return
			case nil:
				successDTO := SuccessResponseDTO[map[domain.ID]domain.MinerInfo]{
					Data:    &inactiveMiners,
					Message: "",
					Time:    time.Now(),
				}
				SendJSON(response, successDTO, http.StatusOK)
				return
			default:
				errorDTO := ErrorResponseDTO{
					Error: err.Error(),
					Time:  time.Now(),
				}
				SendJSON(response, errorDTO, http.StatusInternalServerError)
				return
			}
		}

	} else {
		err, inactiveMiners := h.gameService.GetInactiveMiners()

		switch err {
		case service.GameNotRunningYet, service.GameAlreadyFinished:
			errorDTO := ErrorResponseDTO{
				Error: err.Error(),
				Time:  time.Now(),
			}
			SendJSON(response, errorDTO, http.StatusConflict)
			return
		case nil:
			successDTO := SuccessResponseDTO[map[domain.ID]domain.MinerInfo]{
				Data:    &inactiveMiners,
				Message: "",
				Time:    time.Now(),
			}
			SendJSON(response, successDTO, http.StatusOK)
			return
		default:
			errorDTO := ErrorResponseDTO{
				Error: err.Error(),
				Time:  time.Now(),
			}
			SendJSON(response, errorDTO, http.StatusInternalServerError)
			return
		}
	}
}

/*
pattern:	/miners/active
method:		GET
info:		-

succeed:
  - status code:	200 Ok
  - response body:	JSON active miners + time

failed:
  - status code:	400 Bad Request, 409 Conflict, 500 InternalServerError
  - response body: 	JSON with error + time
*/
func (h *HTTPHandlers) GetActiveMiners(response http.ResponseWriter, request *http.Request) {
	if request.URL.Query().Has("type") {
		minerType := request.URL.Query().Get("type")

		if minerType == "" {
			errorDTO := ErrorResponseDTO{
				Error: "The query must contain valid parameter",
				Time:  time.Now(),
			}
			SendJSON(response, errorDTO, http.StatusBadRequest)
			return
		} else {
			err, activeMiners := h.gameService.GetActiveMinersFilter(domain.MinerType(minerType))

			switch err {
			case service.GameNotRunningYet, service.GameAlreadyFinished:
				errorDTO := ErrorResponseDTO{
					Error: err.Error(),
					Time:  time.Now(),
				}
				SendJSON(response, errorDTO, http.StatusConflict)
				return
			case service.MinerTypeNotFound:
				errorDTO := ErrorResponseDTO{
					Error: err.Error(),
					Time:  time.Now(),
				}
				SendJSON(response, errorDTO, http.StatusBadRequest)
				return
			case nil:
				successDTO := SuccessResponseDTO[map[domain.ID]domain.MinerInfo]{
					Data:    &activeMiners,
					Message: "",
					Time:    time.Now(),
				}
				SendJSON(response, successDTO, http.StatusOK)
				return
			default:
				errorDTO := ErrorResponseDTO{
					Error: err.Error(),
					Time:  time.Now(),
				}
				SendJSON(response, errorDTO, http.StatusInternalServerError)
				return
			}
		}

	} else {
		err, activeMiners := h.gameService.GetActiveMiners()

		switch err {
		case service.GameNotRunningYet, service.GameAlreadyFinished:
			errorDTO := ErrorResponseDTO{
				Error: err.Error(),
				Time:  time.Now(),
			}
			SendJSON(response, errorDTO, http.StatusConflict)
			return
		case nil:
			successDTO := SuccessResponseDTO[map[domain.ID]domain.MinerInfo]{
				Data:    &activeMiners,
				Message: "",
				Time:    time.Now(),
			}
			SendJSON(response, successDTO, http.StatusOK)
			return
		default:
			errorDTO := ErrorResponseDTO{
				Error: err.Error(),
				Time:  time.Now(),
			}
			SendJSON(response, errorDTO, http.StatusInternalServerError)
			return
		}
	}
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
