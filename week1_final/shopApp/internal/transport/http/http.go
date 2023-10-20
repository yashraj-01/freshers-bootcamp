package http

import (
	"github.com/gin-gonic/gin"
)

// Customers represents a type that can provide CRUD operations on customers.
type Customers interface {
	CreateCustomer(c *gin.Context)
	GetCustomerByID(c *gin.Context)
}

// Orders represents a type that can provide CRUD operations on orders.
type Orders interface {
	GetOrders(c *gin.Context)
	CreateOrder(c *gin.Context)
	GetOrderByID(c *gin.Context)
}

// Products represents a type that can provide CRUD operations on products.
type Products interface {
	GetProducts(c *gin.Context)
	CreateProduct(c *gin.Context)
	GetProductByID(c *gin.Context)
	UpdateProduct(c *gin.Context)
}

type Server struct {
	customers Customers
	orders    Orders
	products  Products
}

// New will instantiate a new instance of Server.
func New(c Customers, o Orders, p Products) *Server {
	return &Server{
		customers: c,
		orders:    o,
		products:  p,
	}
}

// AddRoutes will add the routes this server supports to the router.
func (s *Server) AddRoutes(r *gin.Engine) {
	grpCustomers := r.Group("/customer-api")
	{
		grpCustomers.POST("customer", s.customers.CreateCustomer)
		grpCustomers.GET("customer/:id", s.customers.GetCustomerByID)
	}

	grpOrders := r.Group("/order-api")
	{
		grpOrders.POST("order", s.orders.CreateOrder)
		grpOrders.GET("order", s.orders.GetOrders)
		grpOrders.GET("order/:id", s.orders.GetOrderByID)
	}

	grpProducts := r.Group("/product-api")
	{
		grpProducts.POST("product", s.products.CreateProduct)
		grpProducts.GET("product", s.products.GetProducts)
		grpProducts.GET("product/:id", s.products.GetProductByID)
		grpProducts.PATCH("product/:id", s.products.UpdateProduct)
	}
}
