package cmd

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"
)

func SummaryCmd() *cobra.Command {
	listCmd := &cobra.Command{
		Use:   "summary",
		Short: "Show expense summary",
		Long:  "Show a summary of expenses for a given month and year.",
		Run: func(cmd *cobra.Command, args []string) {
			month, _ := cmd.Flags().GetInt64("month")
			year, _ := cmd.Flags().GetInt64("year")

			summaryData(&month, &year)
		},
	}

	return listCmd
}

func summaryData(month *int64, year *int64) {

	fileData, err := loadData()
	if err != nil {
		fmt.Println("Error: File not found.")
		return
	}
	var totalAmount float64

	for _, task := range fileData.ExpenseData {
		if *month != 0 && int64(task.Date.Month()) != *month {
			continue
		}
		if *year != 0 && int64(task.Date.Year()) != *year {
			continue
		}
		totalAmount += task.Amount
	}

	if *month != 0 || *year != 0 {
		fmt.Printf("Total Expenses for %s: $%.2f\n", time.Month(*month).String(), totalAmount)
	} else {
		fmt.Printf("Total Expenses: $%.2f\n", totalAmount)
	}
}
