/*
Copyright © Dave Nash dave@davenash.dev

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// dfICmd represents the dfI command
var dfICmd = &cobra.Command{
	Use:   "dfI",
	Short: "Run `df -i` on the remote host and store the output",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("dfI called")
	},
}

func init() {
	fetchCmd.AddCommand(dfICmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// dfICmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// dfICmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
