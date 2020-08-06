package models

// User is a user in the store.
type User struct {
	Name     string
	LastName string
}

// IUserStore is user store interface.
type IUserStore interface {
	Create(u User)
	List() []User
}

// UserStore deals with user persistence.
type UserStore struct {
	users []User
}

// Create creates a new user in the store.
func (us *UserStore) Create(u User) {
	us.users = append(us.users, u)
}

// List lists all users in the store.
func (us *UserStore) List() []User {
	return us.users
}

// NewUserStore creates a new UserStore.
func NewUserStore() *UserStore {
	return &UserStore{
		users: []User{},
	}
}
