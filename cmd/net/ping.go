/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package net

import (
	"fmt"

	"github.com/spf13/cobra"
)

var	(
	urlPath string
)

// pingCmd represents the ping command
var pingCmd = &cobra.Command{
	Use:   "ping",
	Short: "Ping is a subcommand of the net palette",
	Long: ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("ping called")
	},
}

func init() {
	pingCmd.Flags().StringVarP(&urlPath, "url", "u", "", "the url to ping")
	NetCmd.AddCommand(pingCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// pingCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// pingCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
