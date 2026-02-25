package main

import (
	"fmt"
	"log"

	"github.com/abrshDev/auth-rbac/internal/server"
	"github.com/abrshDev/auth-rbac/internal/user"
)

func main() {
	server.LoadEnv()
	server.ConnectDb()
	app := server.NewApp()
	repo := user.NewRepository(server.DB)

	u := &user.User{
		Username:     "test",
		Email:        "test@example.com",
		PasswordHash: "fakehash",
		Role:         "user",
	}

	err := repo.CreateUser(u)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("User created successfully")
	app.Listen(":3000")

}
