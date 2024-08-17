package main

import (
	"net/http"

	"dailyjournal/pkg/db"
	"dailyjournal/pkg/handlers"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {

    db.InitDB()

    r := chi.NewRouter()
    r.Use(middleware.Logger)
    r.Get("/", func(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte("Hello World!"))
    })

	r.Get("/home", func(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte("Welcome to the Home Page!"))
    })
    r.Post("/users", handlers.CreateUserHandler) // Create User
	r.Get("/users", handlers.GetUsersHandler)    // Get All Users
	
    http.ListenAndServe(":3000", r)
    
}