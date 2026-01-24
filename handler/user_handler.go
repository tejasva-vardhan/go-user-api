package handler

import (
	"encoding/json" // JSON decode/encode ke liye
	"net/http"      // HTTP methods + status codes ke liye
	"strconv"       // string -> int64 convert karne ke liye
	"strings"       // URL se id extract karne ke liye

	"github.com/tejasva-vardhan/go-user-api/model" // User struct
	"github.com/tejasva-vardhan/go-user-api/store" // Repository interface
)

// UserHandler ka kaam: HTTP request ko store (repo) se connect karna
type UserHandler struct {
	Store store.UserRepository // Ab handler interface pe depend karega (DB / memory dono support)
}

// NewUserHandler constructor: store attach karke handler return karta hai
func NewUserHandler(s store.UserRepository) *UserHandler {
	return &UserHandler{Store: s}
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

		// Store se saare users nikaal lo (DB se)
		users, err := h.Store.GetAllUsers(r.Context())
		if err != nil {
			http.Error(w, `{"error":"failed to fetch users"}`, http.StatusInternalServerError)
			return
		} // db error

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
			http.Error(w, `{"error":"invalid json body"}`, http.StatusBadRequest)
			return
		} // invalid json

		// Basic validation (handler level)
		if strings.TrimSpace(input.Name) == "" || strings.TrimSpace(input.Email) == "" {
			http.Error(w, `{"error":"name and email are required"}`, http.StatusBadRequest)
			return
		} // validation fail

		// Store me user create karvao
		createdUser, err := h.Store.CreateUser(r.Context(), input)
		if err != nil {
			// DB insert error => 500 (unique email etc bhi yahi aa sakta hai)
			http.Error(w, `{"error":"failed to create user"}`, http.StatusInternalServerError)
			return
		} // create failed

		// Success => 201 Created
		w.WriteHeader(http.StatusCreated)

		// Created user JSON me return
		json.NewEncoder(w).Encode(createdUser)
		return

	// -----------------------------
	// Any other method => 405
	// -----------------------------
	default:
		http.Error(w, `{"error":"method not allowed"}`, http.StatusMethodNotAllowed)
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
		http.Error(w, `{"error":"user id is required"}`, http.StatusBadRequest)
		return
	} // missing id

	// idStr ko int64 me convert karo
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		// Agar number nahi hai => 400
		http.Error(w, `{"error":"user id must be a number"}`, http.StatusBadRequest)
		return
	} // invalid id

	// Ab method ke basis pe decide karo
	switch r.Method {

	// -----------------------------
	// GET /users/{id}
	// -----------------------------
	case http.MethodGet:

		// Store se user fetch karo
		user, err := h.Store.GetUserByID(r.Context(), id)
		if err != nil {
			// User nahi mila => 404
			http.Error(w, `{"error":"user not found"}`, http.StatusNotFound)
			return
		} // not found / db error

		// User JSON me return
		json.NewEncoder(w).Encode(user)
		return

	// -----------------------------
	// DELETE /users/{id}
	// -----------------------------
	case http.MethodDelete:

		// Store se delete karvao
		err := h.Store.DeleteUser(r.Context(), id)
		if err != nil {
			// User nahi mila => 404
			http.Error(w, `{"error":"user not found"}`, http.StatusNotFound)
			return
		} // delete failed

		// Success => 204 No Content (best practice)
		w.WriteHeader(http.StatusNoContent)
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
			http.Error(w, `{"error":"invalid json body"}`, http.StatusBadRequest)
			return
		} // invalid json

		// Basic validation
		if strings.TrimSpace(input.Name) == "" || strings.TrimSpace(input.Email) == "" {
			http.Error(w, `{"error":"name and email are required"}`, http.StatusBadRequest)
			return
		} // validation fail

		// Store me update karvao
		updatedUser, err := h.Store.UpdateUser(r.Context(), id, input)
		if err != nil {
			// user not found => 404
			http.Error(w, `{"error":"user not found"}`, http.StatusNotFound)
			return
		} // update failed

		// Updated user return
		json.NewEncoder(w).Encode(updatedUser)
		return

	// -----------------------------
	// Any other method => 405
	// -----------------------------
	default:
		http.Error(w, `{"error":"method not allowed"}`, http.StatusMethodNotAllowed)
		return
	}
}
