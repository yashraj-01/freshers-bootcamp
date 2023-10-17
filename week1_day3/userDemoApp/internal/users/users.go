package users

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"yash/userDemoApp/internal/users/model"
	"yash/userDemoApp/internal/users/store"
)

type Users struct {
	store store.Store
}

// New will instantiate a new instance of Users.
func (u *Users) New(s store.Store) *Users {
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
		c.JSON(http.StatusOK, *users)
	}
}

// CreateUser ... Create User
func (u *Users) CreateUser(c *gin.Context) {
	var user *model.User
	c.BindJSON(user)
	err := u.store.CreateUser(user)
	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, *user)
	}
}

// GetUserByID ... Get the user by id
func (u *Users) GetUserByID(c *gin.Context) {
	id := c.Params.ByName("id")
	user, err := u.store.GetUserByID(id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, *user)
	}
}

// UpdateUser ... Update the user information
func (u *Users) UpdateUser(c *gin.Context) {
	id := c.Params.ByName("id")
	user, err := u.store.GetUserByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, *user)
	}
	c.BindJSON(user)
	err = u.store.UpdateUser(user)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, *user)
	}
}

// DeleteUser ... Delete the user
func (u *Users) DeleteUser(c *gin.Context) {
	var user *model.User
	id := c.Params.ByName("id")
	err := u.store.DeleteUser(user, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, gin.H{"id" + id: "is deleted"})
	}
}
