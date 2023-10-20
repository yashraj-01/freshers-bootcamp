package orders

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"yash/shopApp/internal/orders/model"
)

// Store represents a type for storing a user in a database.
type Store interface {
	GetAllOrders() (*[]model.Order, error)
	CreateOrder(order *model.Order) error
	GetOrderByID(id string) (*model.Order, error)
}

type Orders struct {
	store Store
}

// New will instantiate a new instance of Orders.
func New(s Store) *Orders {
	return &Orders{
		store: s,
	}
}

// GetOrders ... Get all orders
func (o *Orders) GetOrders(c *gin.Context) {
	orders, err := o.store.GetAllOrders()
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, orders)
	}
}

// CreateOrders ... Create Order
func (o *Orders) CreateOrders(c *gin.Context) {
	var order model.Order
	c.BindJSON(&order)
	err := o.store.CreateOrder(&order)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, order)
	}
}

// GetOrderByID ... Get the order by id
func (o *Orders) GetOrderByID(c *gin.Context) {
	id := c.Params.ByName("id")
	order, err := o.store.GetOrderByID(id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, order)
	}
}
