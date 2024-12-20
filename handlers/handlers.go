// Manages user interaction for the expense tracker
package handlers

import (
	"bufio"
	"expense_tracker/entities"
	"expense_tracker/services"
	"fmt"
	"os"
	"strings"
	"time"
)

func Start() {
	var expenses []entities.Expense

	for {
		fmt.Println("\n1. Add Expense")
		fmt.Println("2. Save to CSV")
		fmt.Println("3. Analyze Expenses")
		fmt.Println("4. Exit")
		fmt.Print("Choose an option: ")

		var choice int
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			addExpense(&expenses)
		case 2:
			saveToCSV(expenses)
		case 3:
			analyzeExpenses(expenses)
		case 4:
			fmt.Println("Goodbye!")
			return
		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}

func addExpense(expenses *[]entities.Expense) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter date (YYYY-MM-DD): ")
	var dateStr string
	fmt.Scanln(&dateStr)
	date, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		fmt.Println("Invalid date format. Please try again.")
		return
	}

	fmt.Print("Enter description: ")
	description, _ := reader.ReadString('\n')
	description = strings.TrimSpace(description)

	fmt.Print("Enter amount: ")
	var amount float64
	fmt.Scanln(&amount)

	fmt.Print("Enter category: ")
	category, _ := reader.ReadString('\n')
	category = strings.TrimSpace(category)

	*expenses = append(*expenses, entities.Expense{
		Date:        date,
		Description: description,
		Amount:      amount,
		Category:    category,
	})

	fmt.Println("Expense added successfully!")
}

func saveToCSV(expenses []entities.Expense) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter filename (e.g., expenses.csv): ")
	filename, _ := reader.ReadString('\n')
	filename = strings.TrimSpace(filename)

	if err := services.SaveExpensesToCSV(expenses, filename); err != nil {
		fmt.Println("Error saving to CSV:", err)
		return
	}

	fmt.Println("Expenses saved to", filename)
}

func analyzeExpenses(expenses []entities.Expense) {
	totals := services.AnalyzeExpenses(expenses)
	fmt.Println("\nExpense Analysis:")
	for category, total := range totals {
		fmt.Printf("Category: %s, Total: %.2f\n", category, total)
	}
}
