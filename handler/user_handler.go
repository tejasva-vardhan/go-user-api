package handler

import (
	"encoding/json" // JSON decode/encode ke liye
	"net/http"      // HTTP methods + status codes ke liye
	"strconv"       // string -> int convert karne ke liye
	"strings"       // URL se id extract karne ke liye

	"github.com/tejasva-vardhan/go-user-api/model" // User struct
	"github.com/tejasva-vardhan/go-user-api/store" // Store methods
)

// UserHandler ka kaam: HTTP request ko store se connect karna
type UserHandler struct {
	Store *store.UserStore // Store dependency (data yahi handle karega)
}

// NewUserHandler constructor: store attach karke handler return karta hai
func NewUserHandler(s *store.UserStore) *UserHandler {
	return &UserHandler{
		Store: s, // handler ke andar store set
	}
}

// UsersHandler /users route handle karega
// GET  /users  -> list all users
// POST /users  -> create user
func (h *UserHandler) UsersHandler(w http.ResponseWriter, r *http.Request) {

	// Response JSON me bhejna hai
	w.Header().Set("Content-Type", "application/json")

	// Method ke basis pe decide karenge kya karna hai
	switch r.Method {

	// -----------------------------
	// GET /users => List all users
	// -----------------------------
	case http.MethodGet:

		// Store se saare users nikaal lo
		users := h.Store.GetAllUsers()

		// JSON me encode karke response bhej do
		json.NewEncoder(w).Encode(users)
		return

	// -----------------------------
	// POST /users => Create new user
	// -----------------------------
	case http.MethodPost:

		// Input struct jisme JSON decode hoga
		var input model.User

		// Request body JSON ko struct me convert kar rahe hain
		err := json.NewDecoder(r.Body).Decode(&input)
		if err != nil {
			// Invalid JSON => 400
			http.Error(w, "Invalid JSON body", http.StatusBadRequest)
			return
		}

		// Store me user create karvao (validation store karega)
		createdUser, err := h.Store.CreateUser(input)
		if err != nil {
			// Validation fail => 400
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// Success => 201 Created
		w.WriteHeader(http.StatusCreated)

		// Created user JSON me return
		json.NewEncoder(w).Encode(createdUser)
		return

	// -----------------------------
	// Any other method => 405
	// -----------------------------
	default:
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
}

// UserByIDHandler /users/{id} route handle karega
// GET    /users/{id} -> single user
// DELETE /users/{id} -> delete user
// PUT    /users/{id} -> update user
func (h *UserHandler) UserByIDHandler(w http.ResponseWriter, r *http.Request) {

	// Response JSON me bhejna hai
	w.Header().Set("Content-Type", "application/json")

	// URL example: /users/5
	// "/users/" prefix hata ke idStr = "5" nikalna
	idStr := strings.TrimPrefix(r.URL.Path, "/users/")

	// Agar idStr empty hai => /users/ (invalid)
	if idStr == "" {
		http.Error(w, "User ID is required", http.StatusBadRequest)
		return
	}

	// idStr ko int me convert karo
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		// Agar number nahi hai => 400
		http.Error(w, "User ID must be a number", http.StatusBadRequest)
		return
	}

	// Ab method ke basis pe decide karo
	switch r.Method {

	// -----------------------------
	// GET /users/{id}
	// -----------------------------
	case http.MethodGet:

		// Store se user fetch karo
		user, exists := h.Store.GetUserByID(id)
		if !exists {
			// User nahi mila => 404
			http.Error(w, "user not found", http.StatusNotFound)
			return
		}

		// User JSON me return
		json.NewEncoder(w).Encode(user)
		return

	// -----------------------------
	// DELETE /users/{id}
	// -----------------------------
	case http.MethodDelete:

		// Store se delete karvao
		deleted := h.Store.DeleteUserByID(id)
		if !deleted {
			// User nahi mila => 404
			http.Error(w, "user not found", http.StatusNotFound)
			return
		}

		// Success message JSON me
		json.NewEncoder(w).Encode(map[string]string{
			"message": "user deleted",
		})
		return

	// -----------------------------
	// PUT /users/{id}
	// -----------------------------
	case http.MethodPut:

		// Update ke liye input JSON decode karna padega
		var input model.User

		// Body decode
		err := json.NewDecoder(r.Body).Decode(&input)
		if err != nil {
			http.Error(w, "Invalid JSON body", http.StatusBadRequest)
			return
		}

		// Store me update karvao
		updatedUser, exists, err := h.Store.UpdateUserByID(id, input)

		// User exist hi nahi karta => 404
		if !exists {
			http.Error(w, "user not found", http.StatusNotFound)
			return
		}

		// Validation error => 400
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// Updated user return
		json.NewEncoder(w).Encode(updatedUser)
		return

	// -----------------------------
	// Any other method => 405
	// -----------------------------
	default:
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
}
