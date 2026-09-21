package transfer

import (
	"net/http"
	"time"

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
