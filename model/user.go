package model

// User struct ek user ka data represent karta hai
// Ye struct API request, response aur future DB ke liye use hoga
type User struct {

	// ID user ka unique identifier hoga
	ID int `json:"id"`

	// Name user ka naam store karega
	Name string `json:"name"`

	// Email user ka email address store karega
	Email string `json:"email"`
}
