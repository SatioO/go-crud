package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"
	"github.com/satioO/basics/v2/api"
	"github.com/satioO/basics/v2/usecase"
)

func main() {
	router := mux.NewRouter()

	userService := usecase.NewUserService()
	api.RegisterUserHandler(router, userService)

	server := http.Server{
		Handler: router,
		Addr:    ":8080",
	}

	go func() {
		fmt.Println("app started listening on port 8080")
		log.Fatal(server.ListenAndServe())
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, os.Interrupt)

	<-quit

	ctx, shutdown := context.WithTimeout(context.Background(), 5*time.Second)
	defer shutdown()
	server.Shutdown(ctx)
}
