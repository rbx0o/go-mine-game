package transfer

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

// Здесь будет описана логика работы HTTP сервера

type HTTPServer struct {
	httpHandlers *HTTPHandlers
	server       *http.Server
}

/*
NewHTTPServer конструктор для создания объекта типа HTTPServer
*/
func NewHTTPServer(httpHandlers *HTTPHandlers, hostname string, port string) *HTTPServer {
	mux := http.NewServeMux()

	h := &HTTPServer{
		httpHandlers: httpHandlers,
		server: &http.Server{
			Addr:    fmt.Sprintf("%v:%v", hostname, port),
			Handler: mux,
		},
	}

	// endpoints /game/*
	mux.HandleFunc("POST /game/start", h.httpHandlers.StartGame)
	mux.HandleFunc("POST /game/stop", h.httpHandlers.StopGame)
	mux.HandleFunc("GET /game/state", h.httpHandlers.GetState)
	mux.HandleFunc("GET /game/result", h.httpHandlers.GetGameResult)

	// endpoints /enterprise/*
	mux.HandleFunc("GET /enterprise/info", h.httpHandlers.GetIntermediateEnterpriseInfo)

	// endpoints /miner/*
	mux.HandleFunc("POST /miners", h.httpHandlers.HireMiner)
	mux.HandleFunc("GET /miners/info", h.httpHandlers.GetMinersInfo)
	mux.HandleFunc("GET /miners/all", h.httpHandlers.GetAllMiners)
	mux.HandleFunc("GET /miners/inactive", h.httpHandlers.GetInactiveMiners)
	mux.HandleFunc("GET /miners/active", h.httpHandlers.GetActiveMiners)

	// endpoints /equipment/*
	mux.HandleFunc("POST /equipment", h.httpHandlers.BuyEquipment)
	mux.HandleFunc("GET /equipment", h.httpHandlers.GetEquipment)
	mux.HandleFunc("GET /equipment/info", h.httpHandlers.GetEquipmentInfo)

	// endpoints /server/*
	mux.HandleFunc("POST /server/stop", h.httpHandlers.StopServer)

	return h
}

/*
StartServer запускает HTTP сервер
*/
func (h *HTTPServer) StartServer() <-chan error {
	errorCh := make(chan error, 1)

	go func() {
		err := h.server.ListenAndServe()

		if err != nil && err != http.ErrServerClosed {
			errorCh <- err
		}

		close(errorCh)
	}()

	return errorCh
}

/*
StopServer отправляет сигнал на остановку HTTP сервера
*/
func (h *HTTPServer) StopServer() (err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = h.server.Shutdown(ctx)
	return err
}
