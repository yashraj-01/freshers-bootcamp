package store

import (
	"github.com/jinzhu/gorm"
	"yash/userDemoApp/internal/users/model"
)

// Store provides functionality for working with the database.
type Store struct {
	db gorm.DB
}

// New will instantiate a new instance of Store.
func New(db gorm.DB) *Store {
	return &Store{
		db: db,
	}
}

// GetAllUsers Fetch all user data
func (s *Store) GetAllUsers() (*[]model.User, error) {
	var users *[]model.User
	if err := s.db.Find(users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

// CreateUser ... Insert New data
func (s *Store) CreateUser(user *model.User) error {
	if err := s.db.Create(user).Error; err != nil {
		return err
	}
	return nil
}

// GetUserByID ... Fetch only one user by Id
func (s *Store) GetUserByID(id string) (*model.User, error) {
	var user *model.User
	if err := s.db.Where("id = ?", id).First(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

// UpdateUser ... Update user
func (s *Store) UpdateUser(user *model.User) error {
	if err := s.db.Save(user).Error; err != nil {
		return err
	}
	return nil
}

// DeleteUser ... Delete user
func (s *Store) DeleteUser(user *model.User, id string) error {
	if err := s.db.Where("id = ?", id).Delete(user).Error; err != nil {
		return err
	}
	return nil
}
