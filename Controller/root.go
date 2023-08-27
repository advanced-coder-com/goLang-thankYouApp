package cmd

import (
	"fmt"
	"os"
	"thankYou/Db"
	ThankYouModel "thankYou/Model"

	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "thankYou",
	Short: "Generating 'Thank you' in random language",
	Long:  `This application generates 'Thank you' in random language`,
	Args:  cobra.MaximumNArgs(1),

	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		run(args)
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.thankYou.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func run(args []string) {
	if len(args) == 1 {
		var langCode string = "'" + args[0] + "'"
		conditions := []Db.Condition{{Column: "lang_code", Value: langCode, Predicate: "="}}
		results := ThankYouModel.GetList(conditions)
		if len(results) == 0 {
			noResultsPrint()
		}

		prettyPrint(results)

	} else {
		row := ThankYouModel.GetRandom()
		prettyPrint([]ThankYouModel.ThankYou{row})
	}

}

func prettyPrint(rows []ThankYouModel.ThankYou) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Value", "Language", "Lang Code"})

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

func noResultsPrint() {
	fmt.Println("No results for your query")
}
