package utils

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gofiber/fiber/v2"
)

func ListenAndShutdown(app *fiber.App) {

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		if err := app.Listen(":" + os.Getenv("PORT")); err != nil {
			log.Fatalf("Failed to start server: %v", err)
		}
	}()

	sig := <-sigChan
	fmt.Printf("Received signal %s, shutting down server...\n", sig)
	time.Sleep(time.Second * 5)

	if err := app.Server().Shutdown(); err != nil {
		log.Fatalf("Failed to gracefully shutdown server: %v", err)
	}

	log.Println("Server shutdown complete")

}
