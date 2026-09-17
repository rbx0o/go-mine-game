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
  - status code:	201 Created
  - response body:	Server started successfully

failed:
  - status code:	400 BadRequest, 409 Conflict, 500 InternalServerError
  - response body: 	JSON with error + time
*/
func (h *HTTPHandlers) StartGame(response http.ResponseWriter, request *http.Request) {
	err := h.gameService.Start()
	response.Header().Set("Content-Type", "application/json")

	switch err {
	case service.GameAlreadyFinished:
	case service.GameAlreadyRunning:
		ErrorToJson(err, time.Now(), response, http.StatusConflict)
		return
	case nil:
		response.WriteHeader(http.StatusCreated)
		return
	default:
		ErrorToJson(err, time.Now(), response, http.StatusInternalServerError)
	}

}
