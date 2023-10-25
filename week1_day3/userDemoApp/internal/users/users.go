package users

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"yash/userDemoApp/internal/users/model"
)

// Store represents a type for storing a user in a database.
type Store interface {
	GetAllUsers() (*[]model.User, error)
	CreateUser(user *model.User) error
	GetUserByID(id string) (model.User, error)
	UpdateUser(user *model.User) error
	DeleteUser(user *model.User, id string) error
}

type Users struct {
	store Store
}

// New will instantiate a new instance of Users.
func New(s Store) *Users {
	return &Users{
		store: s,
	}
}

// GetUsers ... Get all users
func (u *Users) GetUsers(c *gin.Context) {
	users, err := u.store.GetAllUsers()
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, users)
	}
}

// CreateUser ... Create User
func (u *Users) CreateUser(c *gin.Context) {
	var user model.User
	c.BindJSON(&user)
	err := u.store.CreateUser(&user)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, user)
	}
}

// GetUserByID ... Get the user by id
func (u *Users) GetUserByID(c *gin.Context) {
	id := c.Params.ByName("id")
	user, err := u.store.GetUserByID(id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, user)
	}
}

// UpdateUser ... Update the user information
func (u *Users) UpdateUser(c *gin.Context) {
	id := c.Params.ByName("id")
	user, err := u.store.GetUserByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, user)
	}
	c.BindJSON(&user)
	err = u.store.UpdateUser(&user)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, user)
	}
}

// DeleteUser ... Delete the user
func (u *Users) DeleteUser(c *gin.Context) {
	var user model.User
	id := c.Params.ByName("id")
	err := u.store.DeleteUser(&user, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, gin.H{"id" + id: "is deleted"})
	}
}
