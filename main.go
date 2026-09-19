package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/rbx0o/go-mine-game/service"
	"github.com/rbx0o/go-mine-game/transfer"
)

func main() {
	_ = godotenv.Load()

	game := service.InitGameService()
	fmt.Println("Game service initialized")

	httpHandlers := transfer.NewHTTPHandlers(game)
	fmt.Println("HTTP handlers initialized")

	server := transfer.NewHTTPServer(httpHandlers)
	fmt.Println("HTTP server initialized")

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

	fmt.Printf("Start HTTP server on %v:%v\n", hostname, port)
	if err := server.StartServer(hostname, port); err != nil {
		str := fmt.Sprintf("HTTP server error %v\n", err)
		panic(str)
	}
	fmt.Println("Game Over")
}
