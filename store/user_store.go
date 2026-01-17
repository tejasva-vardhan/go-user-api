package store

import (
	"errors"  // errors.New() se custom error return karte hain
	"strings" // TrimSpace() ke liye
	"sync"

	"github.com/tejasva-vardhan/go-user-api/model"
)

// Step 1: Data (Struct) - "Dukan mein kya-kya hoga?"
// Sabse pehle socho ki aapko kya-kya store karna hai. Bina data ke logic nahi banta.

// Kaise sochein? "Mujhe users chahiye, unki ID manage karni hai, aur concurrency (safety) ka dhayan rakhna hai."

// Result: Aap UserStore struct banaoge. Isme storage (map), counter (nextID), aur lock (mutex) hoga.

// Step 2: Initialization (Constructor) - "Dukan khulegi kaise?"
// Ab aapko ek function chahiye jo is struct ko "zinda" (initialize) kare.

// Kaise sochein? "Jab program shuru ho, toh mujhe ek khali map chahiye aur ID 1 se start honi chahiye."

// Result: Aap NewUserStore() banaoge. Yeh & use karke pointer return karega taaki poora program ek hi store ko use kare.

// Step 3: Actions (Methods) - "Dukan kaam kaise karegi?"
// Ab socho ki user is store ke saath kya-kya karega? (Create, Get, Update, Delete - CRUD).

// Kaise sochein? "Naya user add karne ke liye mujhe store ke andar ka data change karna padega."

// Result: Jab bhi Store ke andar ka data badalna ho, hamesha Pointer Receiver (s *UserStore) wala method banao.
// UserStore users ko memory me store karta hai

type UserStore struct {

	// mutex concurrency ke liye (future safe)
	mu sync.Mutex

	// users map: key = user ID, value = User struct
	users map[int]model.User

	// nextID naya user create karte waqt ID assign karega
	nextID int
}

func NewUserStore() *UserStore{
	return &UserStore{
		users: make(map[int]model.User),
		nextID: 1,
	}
}

func (s *UserStore) CreateUser(user model.User) (model.User, error) {

	// Lock lagate hain taaki map safe rahe (data race na ho)
	s.mu.Lock()

	// Function end hote hi unlock ho jaaye
	defer s.mu.Unlock()

	// -------------------------
	// Basic Validation Start
	// -------------------------

	// Name clean: extra spaces hatao
	user.Name = strings.TrimSpace(user.Name)

	// Email clean: extra spaces hatao
	user.Email = strings.TrimSpace(user.Email)

	// Name empty hai to error
	if user.Name == "" {
		return model.User{}, errors.New("name is required")
	}

	// Email empty hai to error
	if user.Email == "" {
		return model.User{}, errors.New("email is required")
	}

	// -------------------------
	// Basic Validation End
	// -------------------------

	// ID assign karo
	user.ID = s.nextID

	// Map me save karo
	s.users[user.ID] = user

	// nextID increment karo
	s.nextID++

	// created user return
	return user, nil
}
// GetAllUsers saare users list me return karega
func (s *UserStore) GetAllUsers() []model.User {

	// lock for safe read
	s.mu.Lock()
	defer s.mu.Unlock()

	// slice banayi users store karne ke liye
	result := make([]model.User, 0, len(s.users))

	// map se ek ek user uthao
	for _, user := range s.users {
		result = append(result, user)
	}

	// list return
	return result
}
// GetUserByID ek specific user return karta hai by id
func (s *UserStore) GetUserByID(id int) (model.User, bool) {

	// lock => safe read
	s.mu.Lock()
	defer s.mu.Unlock()

	// map se user nikalo
	user, exists := s.users[id]

	// exists true => user mila
	// exists false => user nahi mila
	return user, exists
}
// DeleteUserByID given id wala user delete karta hai
func (s *UserStore) DeleteUserByID(id int) bool {

	// lock => safe write
	s.mu.Lock()
	defer s.mu.Unlock()

	// check if exists
	_, exists := s.users[id]
	if !exists {
		return false // user nahi mila
	}

	// delete from map
	delete(s.users, id)

	return true // delete success
}
// UpdateUserByID given id wale user ko update karta hai
func (s *UserStore) UpdateUserByID(id int, input model.User) (model.User, bool, error) {

	// lock => safe write
	s.mu.Lock()
	defer s.mu.Unlock()

	// check if user exists
	_, exists := s.users[id]
	if !exists {
		return model.User{}, false, nil
	}

	// clean input
	input.Name = strings.TrimSpace(input.Name)
	input.Email = strings.TrimSpace(input.Email)

	// validation
	if input.Name == "" {
		return model.User{}, true, errors.New("name is required")
	}
	if input.Email == "" {
		return model.User{}, true, errors.New("email is required")
	}

	// create updated user (id fixed)
	updated := model.User{
		ID:    id,
		Name:  input.Name,
		Email: input.Email,
	}

	// save back
	s.users[id] = updated

	return updated, true, nil
}






