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

    log.Println("Servidor rodando em http://localhost:8080")
    http.ListenAndServe(":8080", router)
}
