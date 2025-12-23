package main

import (
	"elhazent/expense-tracker/cmd"
	"fmt"
)

func main() {
	rootCmd := cmd.RootCmd()

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
	}
}