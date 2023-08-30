/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	Console "thankYou/Console"
	ThankYouModel "thankYou/Model"

	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add new record",
	Long:  `You can add one more record thi add command. If record for this lang_code exists it will be overrided`,
	Run: func(cmd *cobra.Command, args []string) {
		add()
	},
}

func init() {
	rootCmd.AddCommand(addCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func add() {
	var newRecord ThankYouModel.ThankYou

	fmt.Print("Enter value: ")
	fmt.Scanln(&newRecord.Value)

	fmt.Print("Enter language: ")
	fmt.Scanln(&newRecord.Language)

	fmt.Print("Enter lang code (like en, fr, rus): ")
	fmt.Scanln(&newRecord.Lang_code)

	newRecord.Save()

	fmt.Println("Entered information:")
	Console.PrettyPrint([]ThankYouModel.ThankYou{newRecord})
}
