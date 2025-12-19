package main

import (
	"fmt"
	"net/http"
)

// main function se hi program start hota hai
func main() {

	// http.HandleFunc ka matlab:
	// jab koi "/health" pe request bheje
	// to ye function chale
	http.HandleFunc("/health", healthHandler)

	// fmt.Println sirf terminal me message print karega
	fmt.Println("Server starting on port 8080...")

	// http.ListenAndServe:
	// 1️⃣ ":8080" = kis port pe server chale
	// 2️⃣ nil = default ServeMux use karo
	err := http.ListenAndServe(":8080", nil)

	// Agar server start hone me error aaya
	// to ye line execute hogi
	if err != nil {
		fmt.Println("Server failed to start:", err)
	}
}

// healthHandler ek HTTP handler function hai
// Ye function tab call hota hai
// jab /health pe request aati hai
func healthHandler(w http.ResponseWriter, r *http.Request) {

	// w (ResponseWriter) = jisme hum response likhte hain
	// r (Request) = jo client ne bheja

	// Status code set kar rahe hain: 200 OK
	w.WriteHeader(http.StatusOK)

	// Response body likh rahe hain
	w.Write([]byte("OK"))
}
