package routes

import (
	"database/sql"
	"net/http"

	"booknest-api/handlers"
	"booknest-api/middleware"

	"github.com/go-chi/chi/v5"
	chimiddleware "github.com/go-chi/chi/v5/middleware"
)

func SetupRoutes(db *sql.DB) http.Handler {
	r := chi.NewRouter()

	r.Use(chimiddleware.Logger)
	r.Use(chimiddleware.Recoverer)

	userHandler := &handlers.UserHandler{DB: db}

	r.Post("/register", userHandler.RegisterUser)
	r.Post("/login", userHandler.LoginUser)

	r.Group(func(r chi.Router) {
		r.Use(middleware.AuthMiddleware)
		r.Get("/profile", userHandler.GetUserProfile)
	})

	return r
} 