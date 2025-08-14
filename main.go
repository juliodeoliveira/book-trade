package main

import (
	"book-trade/config"
	"book-trade/routes"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func main() {
	config.LoadEnv()
	config.ConnectDB()

	var router chi.Router = routes.LoadRoutes()

	fs := http.FileServer(http.Dir("./static"))
	router.Handle("/static/*", http.StripPrefix("/static/", fs))

	serverAddress := config.GetEnv("SERVER_ADDRESS", "")
    log.Printf("Servidor rodando em http://%s \n", serverAddress)
    http.ListenAndServe(serverAddress, router)
}
