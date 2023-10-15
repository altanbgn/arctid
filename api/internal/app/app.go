package app

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"arctid/api/internal/config"
	"arctid/api/internal/database"
  "arctid/api/internal/transport/rest"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/helmet/v2"
)

func Run() {
  config.LoadDotenv()
  config := config.Get()

  // Load Database
  database.Connect()

  // Fiber instance
  server := fiber.New(fiber.Config{
    Prefork: false,
    ReadTimeout: config.TIMEOUT,
  })

  // Middlewares
  server.Use(logger.New())
  server.Use(recover.New())
  server.Use(helmet.New())

  // Routes
  v1 := server.Group("/v1")
  rest.LoadRoutes(v1)

  // Signal channel to capture syscalls
  sig := make(chan os.Signal, 1)
  signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT, os.Interrupt)

  go func() {
    <-sig
    log.Println("Shutting down server...")
    _ = server.Shutdown()
  }()

  // Start server
  serverAddr := fmt.Sprintf("%s:%d", config.HOST, config.PORT)
  err := server.Listen(serverAddr)

  if err != nil {
    log.Fatalf("Server not running smh! error: %v", err)
  }
}
