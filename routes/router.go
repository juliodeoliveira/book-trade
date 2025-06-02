package routes

import (
    "github.com/go-chi/chi/v5"
)

func LoadRoutes() chi.Router {
    router := chi.NewRouter()
    BookRoutes(router);

    return router
}