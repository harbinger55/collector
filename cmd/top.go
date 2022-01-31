/*
Copyright © 2022 Dave Nash dave@davenash.dev

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// topCmd represents the top command
var topCmd = &cobra.Command{
	Use:   "top",
	Short: "Run `top -n 1` on the remote host and store the output",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("top called")
	},
}

func init() {
	fetchCmd.AddCommand(topCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// topCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// topCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
