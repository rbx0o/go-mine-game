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

	mux.HandleFunc("POST /game/start", h.httpHandlers.StartGame)
	mux.HandleFunc("GET /game/state", h.httpHandlers.GetState)

	addr := fmt.Sprintf("%v:%v", hostname, port)
	err := http.ListenAndServe(addr, mux)

	if err != nil && !errors.Is(err, http.ErrServerClosed) {
		return err
	}

	return nil
}
