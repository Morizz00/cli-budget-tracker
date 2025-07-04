package cmd

import (
	"fmt"
	"time"

	"github.com/Morizz00/budget-tracker/data"
	"github.com/Morizz00/budget-tracker/model"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new transaction",
	Run: func(cmd *cobra.Command, args []string) {
		amount, _ := cmd.Flags().GetFloat64("amount")
		entryType, _ := cmd.Flags().GetString("type")
		category, _ := cmd.Flags().GetString("category")
		entry := model.Entry{
			Amount:   amount,
			Type:     entryType,
			Category: category,
			Date:     time.Now(),
		}

		err := data.SaveEntry(entry)
		if err != nil {
			fmt.Println("failed to save entry", err)
			return
		}
		fmt.Printf("entry saved:", entry)
	},
}

func init() {
	addCmd.Flags().Float64("amount", 0, "amount of transaction")
	addCmd.Flags().String("type", "", "type:income or expense")
	addCmd.Flags().String("category", "", "category of transaction")

	rootCmd.AddCommand(addCmd)
}
