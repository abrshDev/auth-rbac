package main

import (
	"fmt"
	"log"

	"github.com/abrshDev/auth-rbac/internal/server"
	"github.com/abrshDev/auth-rbac/pkg/utils"
)

func main() {
	server.LoadEnv()
	server.ConnectDb()
	app := server.NewApp()

	fmt.Println("User created successfully")
	s, _ := utils.GenerateToken(1, "user")
	fmt.Println("s", s)
	log.Fatal(app.Listen(":3000"))

}
