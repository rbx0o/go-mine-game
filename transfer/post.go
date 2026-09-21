package transfer

import (
	"errors"
	"net/http"
	"time"

	"github.com/rbx0o/go-mine-game/domain"
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
			Error: err.Error(),
			Time:  time.Now(),
		}
		SendJSON(response, dto, http.StatusInternalServerError)
		return
	}
}

/*
pattern:	/miners
method:		POST
info:		-

succeed:
  - status code:	200 OK
  - response body:	Miner hire successfully + time

failed:
  - status code:	400 Bad Request, 409 Conflict, 500 InternalServerError
  - response body: 	JSON with error + time
*/
func (h *HTTPHandlers) HireMiner(response http.ResponseWriter, request *http.Request) {
	dto := MinerTypeDTO{}
	if err := GetFromJSON(request, &dto); err != nil {
		dtoError := ErrorResponseDTO{
			Error: errors.New("The request body must contain valid JSON").Error(),
			Time:  time.Now(),
		}
		SendJSON(response, dtoError, http.StatusBadRequest)
		return
	}

	if dto.Type == "" {
		dtoError := ErrorResponseDTO{
			Error: errors.New("The field type is required").Error(),
			Time:  time.Now(),
		}
		SendJSON(response, dtoError, http.StatusBadRequest)
		return
	}

	err, miner := h.gameService.HireMiner(domain.MinerType(dto.Type))
	switch err {
	case service.MinerTypeNotFound:
		dtoError := ErrorResponseDTO{
			Error: err.Error(),
			Time:  time.Now(),
		}
		SendJSON(response, dtoError, http.StatusBadRequest)
		return
	case service.NotEnoughCoal:
		dtoError := ErrorResponseDTO{
			Error: err.Error(),
			Time:  time.Now(),
		}
		SendJSON(response, dtoError, http.StatusConflict)
		return
	case service.GameNotRunningYet, service.GameAlreadyFinished:
		dtoError := ErrorResponseDTO{
			Error: err.Error(),
			Time:  time.Now(),
		}
		SendJSON(response, dtoError, http.StatusConflict)
		return
	case nil:
		dtoSuccess := SuccessResponseDTO[domain.MinerInfo]{
			Data:    &miner,
			Message: "The miner has been successfully hired!",
			Time:    time.Now(),
		}
		SendJSON(response, dtoSuccess, http.StatusCreated)
		return
	default:
		dtoError := ErrorResponseDTO{
			Error: err.Error(),
			Time:  time.Now(),
		}
		SendJSON(response, dtoError, http.StatusInternalServerError)
		return
	}
}
