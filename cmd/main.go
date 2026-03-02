package main

import (
	"fmt"
	"log"

	"github.com/abrshDev/auth-rbac/internal/server"
)

func main() {
	server.LoadEnv()
	server.ConnectDb()
	app := server.NewApp()

	fmt.Println("User created successfully")

	log.Fatal(app.Listen(":3000"))

}
