package routes

import (
	"book-trade/controllers"

	"github.com/go-chi/chi/v5"
)

func UserRoutes(route chi.Router) {
	
	route.Route("/user", func(route chi.Router) {
		// Tem que fazer o post pra logar certo
        route.Get("/login", controllers.Login)
        route.Get("/register", controllers.Register)
		route.Get("/verify", controllers.VerifyUser)
		route.Post("/logout", controllers.Logout)
		route.Post("/login", controllers.LoginUser)
		route.Post("/register", controllers.RegisterUser)
	})
}