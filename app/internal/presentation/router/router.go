package router

import (
	"log"

	"github.com/soicchi/book_order_system/internal/infrastructure/postgres/repository"
	"github.com/soicchi/book_order_system/internal/logging"
	customerHandler "github.com/soicchi/book_order_system/internal/presentation/handlers/customers"
	"github.com/soicchi/book_order_system/internal/presentation/middlewares"
	customerUseCase "github.com/soicchi/book_order_system/internal/usecase/customers"

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

// Output the all routes to stdout in local when the server starts
func OutputRoutes(e *echo.Echo) {
	routes := e.Routes()
	for _, route := range routes {
		if route.Method != echo.RouteNotFound {
			log.Printf("%s %s ->\n", route.Method, route.Path)
		}
	}
}
