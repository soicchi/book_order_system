package router

import (
	"log"

	"github.com/soicchi/book_order_system/internal/infrastructure/postgres/repository"
	"github.com/soicchi/book_order_system/internal/logging"
	customerHandler "github.com/soicchi/book_order_system/internal/presentation/handlers/customers"
	orderHandler "github.com/soicchi/book_order_system/internal/presentation/handlers/orders"
	shippingAddressHandler "github.com/soicchi/book_order_system/internal/presentation/handlers/shippingAddresses"
	"github.com/soicchi/book_order_system/internal/presentation/middlewares"
	customerUseCase "github.com/soicchi/book_order_system/internal/usecase/customers"
	orderUseCase "github.com/soicchi/book_order_system/internal/usecase/orders"
	shippingAddressUseCase "github.com/soicchi/book_order_system/internal/usecase/shippingAddresses"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// /api
func NewRouter(e *echo.Echo, logger logging.Logger) {
	e.Pre(middleware.RemoveTrailingSlash())
	base := e.Group("/api")

	// set up common middlewares
	base.Use(middleware.BodyDump(middlewares.CustomBodyDump(logger)))

	v1Router(base, logger)
}

// /api/v1
func v1Router(base *echo.Group, logger logging.Logger) {
	v1 := base.Group("/v1")

	customerRouter(v1, logger)
	shippingAddressRouter(v1, logger)
	orderRouter(v1, logger)
}

// /api/{version}/customers
func customerRouter(version *echo.Group, logger logging.Logger) {
	customersPath := version.Group("/customers")

	// Initialize dependencies
	repo := repository.NewCustomerRepository()
	uc := customerUseCase.NewCustomerUseCase(repo, logger)
	handler := customerHandler.NewCustomerHandler(uc, logger)

	customersPath.POST("/", handler.CreateCustomer)
	customersPath.GET("/:id", handler.FetchCustomer)
}

// /api/{version}/customers/:customer_id/shipping_addresses
func shippingAddressRouter(version *echo.Group, logger logging.Logger) {
	shippingAddressesPath := version.Group("/customers/:customer_id/shipping_addresses")

	// Initialize dependencies
	shippingRepo := repository.NewShippingAddressRepository()
	customerRepo := repository.NewCustomerRepository()
	uc := shippingAddressUseCase.NewShippingAddressUseCase(shippingRepo, customerRepo, logger)
	handler := shippingAddressHandler.NewShippingAddressHandler(uc, logger)

	shippingAddressesPath.POST("/", handler.CreateShippingAddress)
}

// /api/{version}/customers/:customer_id/orders
func orderRouter(version *echo.Group, logger logging.Logger) {
	ordersPath := version.Group("/customers/:customer_id/orders")

	// Initialize dependencies
	orderRepo := repository.NewOrderRepository()
	customerRepo := repository.NewCustomerRepository()
	shippingAddressRepo := repository.NewShippingAddressRepository()
	uc := orderUseCase.NewOrderUseCase(orderRepo, customerRepo, shippingAddressRepo, logger)
	handler := orderHandler.NewOrderHandler(uc, logger)

	ordersPath.POST("/", handler.CreateOrder)
}

// Output the all routes to stdout in local when the server starts
func OutputRoutes(e *echo.Echo) {
	routes := e.Routes()
	for _, route := range routes {
		if route.Method != echo.RouteNotFound {
			log.Printf("%s %s ->\n", route.Method, route.Path)
		}
	}
}
