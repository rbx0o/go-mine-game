package transfer

import (
	"net/http"
	"time"

	"github.com/rbx0o/go-mine-game/service"
)

// Здесь будут описаны обработчики с методом POST

/*
pattern:	/game/start
method:		POST
info:		-

succeed:
  - status code:	200 OK
  - response body:	Game started successfully + time

failed:
  - status code:	409 Conflict, 500 InternalServerError
  - response body: 	JSON with error + time
*/
func (h *HTTPHandlers) StartGame(response http.ResponseWriter, request *http.Request) {
	err := h.gameService.Start()

	switch err {
	case service.GameAlreadyFinished, service.GameAlreadyRunning:
		dto := ErrorResponseDTO{
			Error: err.Error(),
			Time:  time.Now(),
		}
		SendJSON(response, dto, http.StatusConflict)
		return
	case nil:
		dto := SuccessResponseDTO[struct{}]{
			Data:    nil,
			Message: "Game started successfully",
			Time:    time.Now(),
		}
		SendJSON(response, dto, http.StatusOK)
		return
	default:
		dto := ErrorResponseDTO{
			Error: "",
			Time:  time.Now(),
		}
		SendJSON(response, dto, http.StatusInternalServerError)
		return
	}
}
