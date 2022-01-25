/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"calc/webserver"

	"github.com/spf13/cobra"
)

var port string

// htppCmd represents the htpp command
var httpCmd = &cobra.Command{
	Use:   "http",
	Short: "Run http server",
	Long:  `Desc long`,
	Run: func(cmd *cobra.Command, args []string) {
		server := webserver.Server{Port: port}
		server.Serve()
	},
}
