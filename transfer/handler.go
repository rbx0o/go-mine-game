package transfer

import "github.com/rbx0o/go-mine-game/service"

type HTTPHandlers struct {
	gameService *service.GameService
	StopCh      chan struct{}
}

func NewHTTPHandlers(gameService *service.GameService) *HTTPHandlers {
	return &HTTPHandlers{
		gameService: gameService,
		StopCh:      make(chan struct{}),
	}
}
