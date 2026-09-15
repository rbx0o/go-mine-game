package transfer

import "net/http"

// Здесь будут описаны обработчики с методом POST

/*
pattern:	/game/start
method:		POST
info:

succeed:
  - status code:	201 Created
  - response body:	JSON created task

failed:
  - status code:	400 BadRequest, 409 Conflict, 500 InternalServerError
  - response body: 	JSON with error + time
*/
func (h *HTTPHandlers) StartGame(response http.ResponseWriter, request *http.Request) {
	h.gameService.Start()
}
