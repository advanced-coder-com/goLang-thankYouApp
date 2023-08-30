package console

import (
	"os"

	ThankYouModel "thankYou/Model"

	"github.com/olekukonko/tablewriter"
)

func PrettyPrint(rows []ThankYouModel.ThankYou) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Value", "Language", "Lang Code", "To Copy"})

	var data [][]string

	for _, object := range rows {
		data = append(data, object.ToArray())
	}

	// Append data to the table
	for _, v := range data {
		table.Append(v)
	}

	// Set table formatting options
	table.SetBorder(true)
	table.SetCenterSeparator("|")
	table.SetColumnSeparator("|")
	table.SetHeaderAlignment(tablewriter.ALIGN_LEFT)
	table.SetAlignment(tablewriter.ALIGN_LEFT)

	// Render the table
	table.Render()
}
