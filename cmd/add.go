package cmd

import (
	"elhazent/expense-tracker/model"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"time"

	"github.com/spf13/cobra"
)

func AddCmd() *cobra.Command {

	AddCmd := &cobra.Command{
		Use:   "add",
		Short: "Add a new expense",
		Long:  `Add a new expense to your expense tracker with details like, description and amount.`,
		Run: func(cmd *cobra.Command, args []string) {
			// Implementation for adding a new expense
			fmt.Println("Adding a new expense...")
			description, _ := cmd.Flags().GetString("description")
			category, _ := cmd.Flags().GetString("category")
			amount, _ := cmd.Flags().GetFloat64("amount")
			fmt.Printf("Expense Added: %s - $%.2f %s\n", description, amount, category)
			if description == "" || amount <= 0 {
				fmt.Println("Please provide valid description and amount for the expense.")
				return
			}
			addTask(description, amount, &category)
		},
	}

	return AddCmd

}

func addTask(description string, amount float64, category *string) {
	dataStore, err := loadData()
	if err != nil {
		log.Fatalf("Error loading data: %v", err)
	}
	newID := getNextID()
	newData := model.ExpenseModel{
		Id:          newID,
		Description: description,
		Amount:      amount,
		Date:        time.Now(),
		Category:    *category,
	}
	dataStore.ExpenseData = append(dataStore.ExpenseData, newData)

	jsonData, err := json.MarshalIndent(dataStore, "", "  ")
	if err != nil {
		log.Fatalf("Error saving data: %v", err)
	}

	err = ioutil.WriteFile(fileName, jsonData, 0644)
	if err != nil {
		log.Fatalf("Error saving data: %v", err)
	}

	fmt.Printf("Expense added successfully (ID: %d)", newData.Id)
}
