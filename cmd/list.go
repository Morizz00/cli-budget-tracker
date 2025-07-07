package cmd

import (
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/Morizz00/budget-tracker/data"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Display all the entries",
	Run: func(cmd *cobra.Command, args []string) {
		w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
		fmt.Fprintln(w, "Index\tDate\tAmount\tType\tCategory")
		entries, err := data.LoadEntries()
		if err != nil {
			fmt.Printf("Can't load shi")
			return
		}
		if len(entries) == 0 {
			fmt.Println("No entries found.")
			return
		}
		for idx, entry := range entries {
			fmt.Fprintf(w, "%d\t%s\t%.2f\t%s\t%s\n",
				idx,
				entry.Date.Format("2006-01-02"),
				entry.Amount,
				entry.Type,
				entry.Category,
			)
		}
		w.Flush()
	},
}
