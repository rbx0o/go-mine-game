package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/rbx0o/go-mine-game/service"
	"github.com/rbx0o/go-mine-game/transfer"
)

func StopAll(game *service.GameService, server *transfer.HTTPServer) error {
	err, _ := game.StopGame()
	if err != service.GameNotRunningYet &&
		err != service.GameAlreadyFinished &&
		err != nil {
		return err
	}

	err = server.StopServer()
	if err != nil {
		return err
	}

	return nil
}

func main() {
	_ = godotenv.Load()

	game := service.InitGameService()
	fmt.Println("Game service initialized")

	httpHandlers := transfer.NewHTTPHandlers(game)
	fmt.Println("HTTP handlers initialized")

	hostname, ok := os.LookupEnv("HTTP_HOSTNAME")
	if !ok {
		fmt.Println("HTTP_HOSTNAME is required")
		return
	}
	port, ok := os.LookupEnv("HTTP_PORT")
	if !ok {
		fmt.Println("HTTP_PORT is required")
		return
	}

	server := transfer.NewHTTPServer(httpHandlers, hostname, port)
	fmt.Println("HTTP server initialized")

	srvErrCh := server.StartServer()
	fmt.Printf("Start HTTP server on %v:%v\n", hostname, port)
	select {
	case err := <-srvErrCh:
		if err != nil {
			str := fmt.Sprintf("HTTP server error %v\n", err)
			panic(str)
		}
	case <-httpHandlers.StopCh:
		err := StopAll(game, server)
		if err != nil {
			fmt.Printf("Stop game error %v\n", err)
		}
	}

	fmt.Println("Game Over")
}
