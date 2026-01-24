package main

import (
	"github.com/tejasva-vardhan/go-user-api/db"
	"log"
	"net/http"

	// Apna handler package import kar rahe hain
	"github.com/tejasva-vardhan/go-user-api/handler"
	"github.com/tejasva-vardhan/go-user-api/store"
)

func main() {
	database, err := db.OpenSQLite()
	if err != nil {
		log.Fatal(err)
	} // db open failed

	defer database.Close() // cleanup

	err = db.Migrate(database)
	if err != nil {
		log.Fatal(err)
	} // migration failed

	log.Println("DB OK ✅ Migration OK ✅")

	// (1) Store create
	userRepo := store.NewSQLiteUserStore(database)

	// (2) Handler create (store inject)
	userHandler := handler.NewUserHandler(userRepo)

	// (3) Routes register
	http.HandleFunc("/health", handler.HealthHandler)
	http.HandleFunc("/users", userHandler.UsersHandler)
	http.HandleFunc("/users/", userHandler.UserByIDHandler)

	log.Println("Server running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
