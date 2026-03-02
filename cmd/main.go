package main

import (
	"fmt"
	"log"
	"os"

	"github.com/abrshDev/auth-rbac/internal/server"
)

func main() {
	server.LoadEnv()
	server.ConnectDb()

	app := server.NewApp()

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000" // local fallback
	}

	fmt.Println("Server running on port:", port)

	log.Fatal(app.Listen(":" + port))
}
