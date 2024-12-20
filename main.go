package main

import (
	"expense_tracker/handlers"
	"fmt"
)

func main() {
	// Welcome message and start the program
	fmt.Println("Welcome to the Expense Tracker!")
	handlers.Start()
}
