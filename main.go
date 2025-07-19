package main

import (
	"l0s0s/Rarible-API/config"
	"l0s0s/Rarible-API/handler"
	"l0s0s/Rarible-API/rarible"
	"l0s0s/Rarible-API/service"
	"net/http"

	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	config, err := config.ReadConfig()
	if err != nil {
		log.Fatalf("Error reading config: %v", err)
	}

	raribleClient := rarible.NewClient(config.APIKey, config.Referer)
	service := service.NewService(raribleClient)
	handler := handler.NewHandler(service)
	router := gin.Default()
	handler.RegisterRoutes(router)

	srv := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}
}
