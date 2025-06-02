package routes

import (
	"book-trade/controllers"

	"github.com/go-chi/chi/v5"
)

func BookRoutes(route chi.Router) {
    route.Get("/", controllers.Home);
	// route.Get("/send-book", controllers.SendBook);

	route.Route("/books", func(route chi.Router) {
		route.Get("/add", controllers.AddBookPage)
		route.Post("/send", controllers.SendBook)
	})
}
