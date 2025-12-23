package cmd

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
)

func ListCmd() *cobra.Command {
	listCmd := &cobra.Command{
		Use:   "list",
		Short: "List all expenses",
		Long:  "List all recorded expenses in a tabular format.",
		Run: func(cmd *cobra.Command, args []string) {
			category, _ := cmd.Flags().GetString("category")
			if category != "" {
				listFilterData(&category)
				return
			}
			listFilterData(nil)
		},
	}

	return listCmd
}

func listFilterData(category *string) {

	fileData, err := loadData()
	if err != nil {
		fmt.Println("Error: File not found.")
		return
	}
	table := tablewriter.NewWriter(os.Stdout)
	table.Header([]string{"ID", "Description", "Amount", "Category", "Date"})

	for _, task := range fileData.ExpenseData {
		table.Append([]string{
			strconv.Itoa(task.Id),
			task.Description,
			fmt.Sprintf("%.2f", task.Amount),
			task.Category,
			task.Date.Format(time.RFC3339),
		})
	}
	table.Render()
}
