package routes

import (
	"book-trade/controllers"
	"book-trade/middleware"

	"github.com/go-chi/chi/v5"
)

func BookRoutes(route chi.Router) {
    route.Get("/", controllers.Home);

	route.Route("/books", func(route chi.Router) {
		route.With(middleware.RequireAuth).Get("/add", controllers.AddBookPage)
		route.With(middleware.RequireAuth).Post("/send", controllers.SendBook)
		route.With(middleware.RequireAuth).Get("/my-books", controllers.UserBooksPage)
		route.With(middleware.RequireAuth).Post("/delete", controllers.DeleteBook)
	})
}
