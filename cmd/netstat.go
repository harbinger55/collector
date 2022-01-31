/*
Copyright © 2022 Dave Nash dave@davenash.dev

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// netstatCmd represents the netstat command
var netstatCmd = &cobra.Command{
	Use:   "netstat",
	Short: "Run `netstat -na ` on the remote host and store the output",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("netstat called")
	},
}

func init() {
	fetchCmd.AddCommand(netstatCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// netstatCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// netstatCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
