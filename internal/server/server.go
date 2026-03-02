package server

import (
	"fmt"
	"log"
	"os"

	"github.com/abrshDev/auth-rbac/internal/auth"
	"github.com/abrshDev/auth-rbac/internal/middleware"
	"github.com/abrshDev/auth-rbac/internal/user"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func LoadEnv() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using environment variables")
	}

}

func ConnectDb() {
	var err error
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)
	//	log.Println("CONNECTION:", dsn)
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	err = DB.AutoMigrate(&user.User{})
	if err != nil {
		log.Fatal("Failed to migrate database:", err)
	}
	fmt.Println("Database connected successfully", DB)
}
func registeroutes(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Auth RBAC Fiber API is running!")
	})
	userRepo := user.NewRepository(DB)

	authHandler := auth.NewAuthHandler(userRepo)

	// Group routes
	api := app.Group("/api")
	authGroup := api.Group("/auth")

	// Register endpoint
	authGroup.Post("/register", authHandler.Register)
	authGroup.Post("/login", authHandler.Login)
	api.Get("/profile",
		middleware.Protected(),
		func(c *fiber.Ctx) error {
			return c.JSON(fiber.Map{
				"user_id": c.Locals("user_id"),
				"role":    c.Locals("role"),
			})
		},
	)

}

func NewApp() *fiber.App {
	app := fiber.New()
	registeroutes(app)
	return app
}
