package cmd

import (
	"encoding/csv"
	"fmt"
	"os"

	"github.com/Morizz00/budget-tracker/data"
	"github.com/spf13/cobra"
)

var exportFile string
var expCmd = &cobra.Command{
	Use:   "export",
	Short: "Can export your CSV files through this",
	Run: func(cmd *cobra.Command, args []string) {
		entries, err := data.LoadEntries()
		if err != nil {
			fmt.Printf("cant load shi")
			return
		}
		f, err := os.Create(exportFile)
		if err != nil {
			fmt.Printf("couldn't create shi")
			return
		}
		writer := csv.NewWriter(f)
		defer f.Close()

		writer.Write([]string{"Date", "Amount", "Type", "Category"})
		if err := writer.Error(); err != nil {
			fmt.Printf("Error writing CSV: %v\n", err)
			return
		}

		for _, entry := range entries {
			writer.Write([]string{
				entry.Date.Format("2006-01-02"),
				fmt.Sprintf("%.2f", entry.Amount),
				entry.Type,
				entry.Category,
			})
		}
		writer.Flush()
		fmt.Printf("Exported %d entries to %s\n", len(entries), exportFile)
	},
}

func init() {
	expCmd.Flags().StringVarP(&exportFile, "output", "o", "entries.csv", "Output CSV file path")
	rootCmd.AddCommand(expCmd)

}
