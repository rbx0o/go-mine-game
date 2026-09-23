package transfer

import (
	"errors"
	"fmt"
	"net/http"
)

// Здесь будет описана логика работы HTTP сервера

type HTTPServer struct {
	httpHandlers *HTTPHandlers
}

/*
NewHTTPServer конструктор для создания объекта типа HTTPServer
*/
func NewHTTPServer(httpHandlers *HTTPHandlers) *HTTPServer {
	return &HTTPServer{
		httpHandlers: httpHandlers,
	}
}

/*
StartServer запускает HTTP сервер
*/
func (h *HTTPServer) StartServer(hostname string, port string) error {
	mux := http.NewServeMux()

	// endpoints /game/*
	mux.HandleFunc("POST /game/start", h.httpHandlers.StartGame)
	mux.HandleFunc("POST /game/stop", h.httpHandlers.StopGame)
	mux.HandleFunc("GET /game/state", h.httpHandlers.GetState)

	// endpoints /enterprise/*
	mux.HandleFunc("GET /enterprise/info", h.httpHandlers.GetIntermediateEnterpriseInfo)

	// endpoints /miner/*
	mux.HandleFunc("POST /miners", h.httpHandlers.HireMiner)
	mux.HandleFunc("GET /miners/info", h.httpHandlers.GetMinersInfo)

	// endpoints /equipment/*
	mux.HandleFunc("POST /equipment", h.httpHandlers.BuqEquipment)

	addr := fmt.Sprintf("%v:%v", hostname, port)
	err := http.ListenAndServe(addr, mux)

	if err != nil && !errors.Is(err, http.ErrServerClosed) {
		return err
	}

	return nil
}
