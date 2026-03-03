package main

import (
	"log"
	"os"

	"github.com/abrshDev/auth-rbac/config"
	"github.com/abrshDev/auth-rbac/internal/server"
)

func main() {
	config.LoadEnv()

	db, err := config.ConnectDb()
	if err != nil {
		log.Fatal(err)
	}

	app := server.NewApp(db)

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	log.Fatal(app.Listen(":" + port))
}
