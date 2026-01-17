package handler

import (
	"encoding/json" // JSON decode/encode ke liye
	"net/http"      // HTTP methods + status codes
    "strconv" // string to int convert

	"strings" // string manipulation
	"github.com/tejasva-vardhan/go-user-api/model"
	"github.com/tejasva-vardhan/go-user-api/store"
)

// UserHandler HTTP layer ko store layer se connect karta hai
type UserHandler struct {
	Store *store.UserStore // store dependency
}

// NewUserHandler handler ka constructor hai
func NewUserHandler(s *store.UserStore) *UserHandler {
	return &UserHandler{
		Store: s, // store attach
	}
}
func (h *UserHandler) UsersHandler(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Content-Type","application/json")

	switch r.Method{
	case http.MethodGet:
		users:=h.Store.GetAllUsers()
		json.NewEncoder(w).Encode(users)

	case http.MethodPost:
		// 3) Request body JSON decode
	var input model.User // yaha input aayega (id empty hoga)

	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		http.Error(w, "Invalid JSON body", http.StatusBadRequest)
		return
	}

	// 4) Store se create karvao
	createdUser, err := h.Store.CreateUser(input)
	if err != nil {
		// validation fail => 400
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// 5) Success => 201 Created
	w.WriteHeader(http.StatusCreated)

	// 6) Created user JSON me return
	json.NewEncoder(w).Encode(createdUser)
	return
		default:
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

}
// UserByIDHandler GET /users/{id} handle karega
func (h *UserHandler) UserByIDHandler(w http.ResponseWriter, r *http.Request) {

	// sirf GET allow
	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	// JSON response
	w.Header().Set("Content-Type", "application/json")

	// URL example: /users/5
	// hume id part chahiye => "5"
	idStr := strings.TrimPrefix(r.URL.Path, "/users/")

	// idStr empty hua => matlab path "/users/" aaya (invalid for this handler)
	if idStr == "" {
		http.Error(w, "User ID is required", http.StatusBadRequest)
		return
	}

	// string -> int convert
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "User ID must be a number", http.StatusBadRequest)
		return
	}

	// store se user fetch
	user, exists := h.Store.GetUserByID(id)
	if !exists {
		http.Error(w, "user not found", http.StatusNotFound)
		return
	}

	// success => 200 + user JSON
	json.NewEncoder(w).Encode(user)
}


