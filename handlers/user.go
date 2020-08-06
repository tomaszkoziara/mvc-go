package handlers

import (
	"fmt"
	"net/http"

	"github.com/tomaszkoziara/mvc-go/models"
)

// IUser is the interface for user handler.
type IUser interface {
	Create(w http.ResponseWriter, r *http.Request)
	List(w http.ResponseWriter, r *http.Request)
}

// User is the user handler for all methods.
type User struct {
	userStore models.IUserStore
}

// Create handlers user creation request.
func (u *User) Create(w http.ResponseWriter, r *http.Request) {
	n := r.FormValue("name")
	ln := r.FormValue("lastname")

	u.userStore.Create(models.User{Name: n, LastName: ln})

	w.WriteHeader(http.StatusCreated)
}

// List handlers user list request.
func (u *User) List(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(fmt.Sprintf("%v", u.userStore.List())))
}

// NewUserHandler creates a new userHandler.
func NewUserHandler(userStore models.IUserStore) *User {
	return &User{
		userStore: userStore,
	}
}
