package main

import (
	"log"
	"net/http"

	// Apna handler package import kar rahe hain
	"github.com/tejasva-vardhan/go-user-api/handler"
	"github.com/tejasva-vardhan/go-user-api/store"
)

func main() {

	// (1) Store create
	userStore := store.NewUserStore()

	// (2) Handler create (store inject)
	userHandler := handler.NewUserHandler(userStore)

	// (3) Routes register
	http.HandleFunc("/health", handler.HealthHandler)
	http.HandleFunc("/users", userHandler.UsersHandler)
    http.HandleFunc("/users/", userHandler.UserByIDHandler)


	log.Println("Server running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
