package store

import (
	"gorm.io/gorm"
	"yash/shopApp/internal/orders/model"
)

// Store provides functionality for working with the database.
type Store struct {
	db *gorm.DB
}

// New will instantiate a new instance of Store.
func New(db *gorm.DB) *Store {
	db.AutoMigrate(&model.Order{})

	return &Store{
		db: db,
	}
}

// GetAllOrders ... Fetch all orders data
func (s *Store) GetAllOrders() (*[]model.Order, error) {
	var orders []model.Order
	if err := s.db.Find(&orders).Error; err != nil {
		return nil, err
	}
	return &orders, nil
}

// CreateOrder ... Insert New data
func (s *Store) CreateOrder(order *model.Order) error {
	if err := s.db.Create(order).Error; err != nil {
		return err
	}
	return nil
}

// GetOrderByID ... Fetch only one order by Id
func (s *Store) GetOrderByID(id string) (*model.Order, error) {
	var order model.Order
	if err := s.db.Where("id = ?", id).First(&order).Error; err != nil {
		return nil, err
	}
	return &order, nil
}
