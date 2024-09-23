package main

import (
	"fmt"

	"github.com/soicchi/book_order_system/infrastructure/postgres"
)

func main() {
	// Database initialization
	postgres.Initialize()
	fmt.Println("Database initialized")
}
