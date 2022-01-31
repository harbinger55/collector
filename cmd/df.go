/*
Copyright © 2022 Dave Nash dave@davenash.dev

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// dfCmd represents the df command
var dfCmd = &cobra.Command{
	Use:   "df",
	Short: "Run `df -h` on the remote host and store the output",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("df called")
	},
}

func init() {
	fetchCmd.AddCommand(dfCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// dfCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// dfCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
