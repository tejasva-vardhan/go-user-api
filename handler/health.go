package handler

import (
	"net/http"
)

// HealthHandler ek HTTP handler function hai
// Ye server ke health check ke liye use hota hai
func HealthHandler(w http.ResponseWriter, r *http.Request) {

	// Status code set kar rahe hain: 200 OK
	w.WriteHeader(http.StatusOK)

	// Response body bhej rahe hain
	w.Write([]byte("OK"))
}
