package main

import (
	"os"
	"server"
)

func main() {
	PORT := os.Getenv("PORT")
	if len(PORT) == 0 {
		PORT = "8080"
	}
	server := service.NewServer()
	server.Run(":" + PORT)
}
