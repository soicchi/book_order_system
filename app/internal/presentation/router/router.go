package router

import (
	"log"

	"github.com/soicchi/book_order_system/internal/infrastructure/postgres/repository"
	"github.com/soicchi/book_order_system/internal/logging"
	booksHandler "github.com/soicchi/book_order_system/internal/presentation/handlers/books"
	usersHandler "github.com/soicchi/book_order_system/internal/presentation/handlers/users"
	"github.com/soicchi/book_order_system/internal/presentation/middlewares"
	booksUseCase "github.com/soicchi/book_order_system/internal/usecase/books"
	usersUseCase "github.com/soicchi/book_order_system/internal/usecase/users"

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

	usersRouter(v1, logger)
	booksRouter(v1, logger)
}

func usersRouter(version *echo.Group, logger logging.Logger) {
	users := version.Group("/users")

	// set up dependencies
	userRepo := repository.NewUserRepository()
	useCase := usersUseCase.NewUseCase(userRepo, logger)
	handler := usersHandler.NewHandler(useCase, logger)

	// routes
	users.POST("", handler.CreateUser)
	users.GET("", handler.ListUsers)
	users.GET("/:user_id", handler.RetrieveUser)
	users.PUT("/:user_id", handler.UpdateUser)
	users.DELETE("/:user_id", handler.DeleteUser)
}

func booksRouter(version *echo.Group, logger logging.Logger) {
	books := version.Group("/books")

	// set up dependencies
	bookRepo := repository.NewBookRepository()
	useCase := booksUseCase.NewUseCase(bookRepo, logger)
	handler := booksHandler.NewHandler(useCase, logger)

	// routes
	books.POST("", handler.CreateBook)
	books.GET("", handler.ListBooks)
	books.GET("/:book_id", handler.RetrieveBook)
	books.PUT("/:book_id", handler.UpdateBook)
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
