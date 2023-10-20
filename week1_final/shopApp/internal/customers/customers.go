package customers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"yash/shopApp/internal/customers/model"
)

// Store represents a type for storing a user in a database.
type Store interface {
	CreateCustomer(customer *model.Customer) error
	GetCustomerByID(id string) (*model.Customer, error)
}

type Customers struct {
	store Store
}

// New will instantiate a new instance of Customers.
func New(s Store) *Customers {
	return &Customers{
		store: s,
	}
}

// CreateCustomers ... Create Customer
func (c *Customers) CreateCustomers(context *gin.Context) {
	var customer model.Customer
	context.BindJSON(&customer)
	err := c.store.CreateCustomer(&customer)
	if err != nil {
		context.AbortWithStatus(http.StatusNotFound)
	} else {
		context.JSON(http.StatusOK, customer)
	}
}

// GetCustomerByID ... Get the customer by id
func (c *Customers) GetCustomerByID(context *gin.Context) {
	id := context.Params.ByName("id")
	customer, err := c.store.GetCustomerByID(id)
	if err != nil {
		context.AbortWithStatus(http.StatusNotFound)
	} else {
		context.JSON(http.StatusOK, customer)
	}
}
