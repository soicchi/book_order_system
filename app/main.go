package main

import "github.com/soicchi/book_order_system/cmd"

// This main function is the all entry point of the application like starting the server,
// running the migration process etc.
func main() {
	cmd.Execute()
}
