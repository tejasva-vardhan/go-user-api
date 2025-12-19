package store

import (
	"sync"

	"github.com/tejasva-vardhan/go-user-api/model"
)

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


