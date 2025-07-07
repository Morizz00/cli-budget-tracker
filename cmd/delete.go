package cmd

import (
	"fmt"

	"github.com/Morizz00/budget-tracker/data"
	"github.com/spf13/cobra"
)

var deleteIndex int

var delCmd = &cobra.Command{
	Use:   "delete",
	Short: "Deletes an entry from the list",
	Run: func(cmd *cobra.Command, args []string) {
		entries, err := data.LoadEntries()
		if err != nil {
			fmt.Printf("can't load shi")
			return
		}
		if deleteIndex < 0 {
			fmt.Printf("please provide a valid index")
			return
		}
		if deleteIndex >= len(entries) {
			fmt.Printf("index out of range")
			return
		}
		deleted := entries[deleteIndex]
		entries = append(entries[:deleteIndex], entries[deleteIndex+1:]...)

		err = data.SaveEntries(entries)
		if err != nil {
			fmt.Printf("failed to save updated entries:%v\n", err)
			return
		}
		fmt.Printf("Deleted entry:#%d: %.2f %s (%s)\n", deleteIndex, deleted.Amount, deleted.Type, deleted.Category)
	},
}

func init() {
	delCmd.Flags().IntVarP(&deleteIndex, "index", "i", -1, "Index of the entry to delete")
	rootCmd.AddCommand(delCmd)
}
