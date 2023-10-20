package store

import (
	"gorm.io/gorm"
	"yash/shopApp/internal/customers/model"
)

// Store provides functionality for working with the database.
type Store struct {
	db *gorm.DB
}

// New will instantiate a new instance of Store.
func New(db *gorm.DB) *Store {
	db.AutoMigrate(&model.Customer{})

	return &Store{
		db: db,
	}
}

// CreateCustomer ... Insert New data
func (s *Store) CreateCustomer(customer *model.Customer) error {
	if err := s.db.Create(customer).Error; err != nil {
		return err
	}
	return nil
}

// GetCustomerByID ... Fetch only one customer by Id
func (s *Store) GetCustomerByID(id string) (*model.Customer, error) {
	var customer model.Customer
	if err := s.db.Where("id = ?", id).First(&customer).Error; err != nil {
		return nil, err
	}
	return &customer, nil
}
