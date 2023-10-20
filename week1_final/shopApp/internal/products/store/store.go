package store

import (
	"gorm.io/gorm"
	"yash/shopApp/internal/products/model"
)

// Store provides functionality for working with the database.
type Store struct {
	db *gorm.DB
}

// New will instantiate a new instance of Store.
func New(db *gorm.DB) *Store {
	db.AutoMigrate(&model.Product{})

	return &Store{
		db: db,
	}
}

// GetAllProducts ... Fetch all products data
func (s *Store) GetAllProducts() (*[]model.Product, error) {
	var products []model.Product
	if err := s.db.Find(&products).Error; err != nil {
		return nil, err
	}
	return &products, nil
}

// CreateProduct ... Insert New data
func (s *Store) CreateProduct(product *model.Product) error {
	if err := s.db.Create(product).Error; err != nil {
		return err
	}
	return nil
}

// GetProductByID ... Fetch only one product by Id
func (s *Store) GetProductByID(id string) (*model.Product, error) {
	var product model.Product
	if err := s.db.Where("id = ?", id).First(&product).Error; err != nil {
		return nil, err
	}
	return &product, nil
}

// UpdateProduct ... Update product
func (s *Store) UpdateProduct(product *model.Product) error {
	s.db.Model(&product).Updates(product)
	if err := s.db.Save(product).Error; err != nil {
		return err
	}
	return nil
}
