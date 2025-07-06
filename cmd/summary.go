package cmd

import (
	"fmt"

	"github.com/Morizz00/budget-tracker/data"
	"github.com/spf13/cobra"
)

var summaryCmd = &cobra.Command{
	Use:   "summary",
	Short: "View total income,expenses and balance",
	Run: func(cmd *cobra.Command, args []string) {
		entries, err := data.LoadEntries()
		if err != nil {
			fmt.Printf("failed to load entries")
			return
		}
		var income float64
		var expense float64

		for _, entry := range entries {
			if entry.Type == "income" {
				income += entry.Amount
			} else if entry.Type == "expense" {
				expense += entry.Amount
			}
		}
		balance := income - expense
		fmt.Printf("Total income:%.2f\n", income)
		fmt.Printf("Total expenses:%.2f\n", expense)
		fmt.Printf("balance:%.2f\n", balance)
	},
}
