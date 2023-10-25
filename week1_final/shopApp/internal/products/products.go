package products

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"yash/shopApp/internal/products/model"
)

// Store represents a type for storing a product in a database.
type Store interface {
	GetAllProducts() (*[]model.Product, error)
	CreateProduct(product *model.Product) error
	GetProductByID(id string) (*model.Product, error)
	UpdateProduct(product *model.Product) error
}

type Products struct {
	store Store
}

// New will instantiate a new instance of Products.
func New(s Store) *Products {
	return &Products{
		store: s,
	}
}

// GetProducts ... Get all products
func (p *Products) GetProducts(c *gin.Context) {
	products, err := p.store.GetAllProducts()
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
	} else {
		c.JSON(http.StatusOK, products)
	}
}

// CreateProduct ... Create Product
func (p *Products) CreateProduct(c *gin.Context) {
	var product model.Product
	c.BindJSON(&product)
	err := p.store.CreateProduct(&product)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
	} else {
		c.JSON(http.StatusOK, product)
	}
}

// GetProductByID ... Get the product by id
func (p *Products) GetProductByID(c *gin.Context) {
	id := c.Params.ByName("id")
	product, err := p.store.GetProductByID(id)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
	} else {
		c.JSON(http.StatusOK, product)
	}
}

// UpdateProduct ... Update the product information
func (p *Products) UpdateProduct(c *gin.Context) {
	id := c.Params.ByName("id")
	product, err := p.store.GetProductByID(id)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
	} else if product == nil {
		c.JSON(http.StatusNotFound, product)
	} else {
		c.BindJSON(&product)
		err = p.store.UpdateProduct(product)
		if err != nil {
			c.AbortWithStatus(http.StatusInternalServerError)
		} else {
			c.JSON(http.StatusOK, product)
		}
	}
}
