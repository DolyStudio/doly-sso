package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/joho/godotenv"
	"github.com/zkfmapf123/pdf-bot/handlers"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/swagger"
)

var (
	DEFAULT_PORT = os.Getenv("PORT")
	APP_NAME     = os.Getenv("APP_NAME")
	ALLOW_ORIGIN = os.Getenv("ALLOW_ORIGIN")
)

// @title			example
// @version		1.0
// @description	This is a sample swagger for Fiber
// @termsOfService	http://swagger.io/terms/
// @contact.name	API Support
// @contact.email	fiber@swagger.io
// @license.name	Apache 2.0
// @license.url	http://www.apache.org/licenses/LICENSE-2.0.html
// @host			localhost:3000
// @BasePath		/
func main() {
	// dotenv
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	app := fiber.New(fiber.Config{
		Prefork:       true,
		CaseSensitive: true,
		StrictRouting: true,
		ServerHeader:  "fiber",
		AppName:       APP_NAME,
	})

	app.Use(logger.New())

	app.Use(cors.New(cors.Config{
		AllowOrigins: ALLOW_ORIGIN,
		AllowMethods: "OPTIONS,GET,POST,DELETE",
		AllowHeaders: "Origin, Content-Type, Accept",
		// AllowCredentials: true,
	}))

	app.Get("/ping", handlers.PingHandlers)
	app.Get("/swagger/*", swagger.HandlerDefault)

	// google
	app.Get("/auth/google/login", handlers.GoogleAuthLogin)
	app.Get("/auth/google/callback", handlers.GoogleAuthCallback)
	// app.Get("/auth/google/logout")

	// naver
	// app.Get("/auth/naver")

	// kakao
	// app.Get("/auth/kakao")

	port := os.Getenv("PORT")
	if port == "" {
		port = DEFAULT_PORT
	}

	go func() {
		if err := app.Listen(":" + port); err != nil {
			log.Printf("Server error: %v\n", err)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	<-c
	log.Println("Shutting down server...")
	app.Shutdown()
}
