// Contains logic for managing expenses and writing to CSV
package services

import (
	"encoding/csv"
	"expense_tracker/entities"
	"fmt"
	"os"
)

// SaveExpensesToCSV writes a list of expenses to a CSV file
func SaveExpensesToCSV(expenses []entities.Expense, filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("could not create file: %w", err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Write CSV headers
	headers := []string{"Date", "Description", "Amount", "Category"}
	if err := writer.Write(headers); err != nil {
		return fmt.Errorf("could not write headers: %w", err)
	}

	// Write expense records
	for _, expense := range expenses {
		record := []string{
			expense.Date.Format("2006-01-02"),
			expense.Description,
			fmt.Sprintf("%.2f", expense.Amount),
			expense.Category,
		}
		if err := writer.Write(record); err != nil {
			return fmt.Errorf("could not write record: %w", err)
		}
	}

	return nil
}

// AnalyzeExpenses calculates total and category-wise expenses
func AnalyzeExpenses(expenses []entities.Expense) map[string]float64 {
	categoryTotals := make(map[string]float64)
	for _, expense := range expenses {
		categoryTotals[expense.Category] += expense.Amount
	}
	return categoryTotals
}
