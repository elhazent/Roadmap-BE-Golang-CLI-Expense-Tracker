package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"elhazent/expense-tracker/model"

	"github.com/spf13/cobra"
)

const fileName = "expense.json"

func RootCmd() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   "expense-tracker",
		Short: "A simple expense tracking application",
		Long:  `Expense Tracker is a CLI application to help you track your expenses efficiently.`,
		Run: func(cmd *cobra.Command, args []string) {
			// Default action when no subcommands are provided
			fmt.Println("Welcome to Expense Trancker CLI Application.....")
		},
	}

	rootCmd.AddCommand(AddCmd())
	rootCmd.AddCommand(ListCmd())
	rootCmd.AddCommand(SummaryCmd())
	rootCmd.PersistentFlags().StringP("description", "D", "", "Description of the expense")
	rootCmd.PersistentFlags().Float64P("amount", "A", 0.0, "Amount of the expense")
	rootCmd.PersistentFlags().StringP("category", "C", "", "Category of the expense (Optional)")
	rootCmd.PersistentFlags().Int64P("month", "M", 0, "Month of the expense (Optional)")
	rootCmd.PersistentFlags().Int64P("year", "Y", 0, "Year of the expense (Optional)")

	return rootCmd
}

func loadData() (model.ExpenseData, error) {
	dataStore := model.ExpenseData{}
	file, err := ioutil.ReadFile(fileName)
	if err != nil {
		if os.IsNotExist(err) {
			return dataStore, nil
		}
		return dataStore, err
	}

	err = json.Unmarshal(file, &dataStore)
	if err != nil {
		return dataStore, err
	}

	return dataStore, nil
}

func getNextID() int {

	loadData, err := loadData()
	if err != nil {
		fmt.Println("Error loading data:", err)
		return 0
	}
	maxID := 0
	for _, record := range loadData.ExpenseData {
		if record.Id > maxID {
			maxID = record.Id
		}
	}
	return maxID + 1
}
