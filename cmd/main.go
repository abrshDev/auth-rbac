package main

import (
	"github.com/abrshDev/auth-rbac/internal/server"
)

func main() {
	server.LoadEnv()
	server.ConnectDb()
	app := server.NewApp()

	app.Listen(":3000")

}
