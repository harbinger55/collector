/*
Copyright © 2022 Dave Nash dave@davenash.dev

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// lsofCmd represents the lsof command
var lsofCmd = &cobra.Command{
	Use:   "lsof",
	Short: "Run `lsof -a` on the remote host and store the output",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("lsof called")
	},
}

func init() {
	fetchCmd.AddCommand(lsofCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// lsofCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// lsofCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
