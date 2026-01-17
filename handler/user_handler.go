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
// UserByIDHandler GET /users/{id} and DELETE /users/{id} handle karega
func (h *UserHandler) UserByIDHandler(w http.ResponseWriter, r *http.Request) {

	// JSON response
	w.Header().Set("Content-Type", "application/json")

	// URL example: /users/5
	// id part nikalna => "5"
	idStr := strings.TrimPrefix(r.URL.Path, "/users/")

	// empty id => invalid
	if idStr == "" {
		http.Error(w, "User ID is required", http.StatusBadRequest)
		return
	}

	// string -> int
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "User ID must be a number", http.StatusBadRequest)
		return
	}

	// METHOD SWITCH
	switch r.Method {

	// GET /users/{id}
	case http.MethodGet:
		user, exists := h.Store.GetUserByID(id)
		if !exists {
			http.Error(w, "user not found", http.StatusNotFound)
			return
		}

		json.NewEncoder(w).Encode(user)
		return

	// DELETE /users/{id}
	case http.MethodDelete:
		deleted := h.Store.DeleteUserByID(id)
		if !deleted {
			http.Error(w, "user not found", http.StatusNotFound)
			return
		}

		// success response
		json.NewEncoder(w).Encode(map[string]string{
			"message": "user deleted",
		})
		return

	// other methods not allowed
	default:
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
}



