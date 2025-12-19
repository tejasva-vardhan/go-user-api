package main

import (
	"fmt"
	"net/http"

	// Apna handler package import kar rahe hain
	"github.com/tejasva-vardhan/go-user-api/handler"
)

func main() {

	// /health endpoint ko handler package ke function se connect kar rahe hain
	http.HandleFunc("/health", handler.HealthHandler)

	fmt.Println("Server starting on port 8080...")

	// Server start kar rahe hain
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Server failed to start:", err)
	}
}
