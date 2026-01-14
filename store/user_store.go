package store

import (
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

func (s *UserStore) CreateUser(user model.User) model.User {

	// Lock lagate hain taaki data race na ho
	s.mu.Lock()
	defer s.mu.Unlock()

	// User ko ID assign kar rahe hain
	user.ID = s.nextID

	// Map me user store kar rahe hain
	s.users[s.nextID] = user

	// nextID increment kar rahe hain
	s.nextID++

	// Created user return kar rahe hain
	return user
}


