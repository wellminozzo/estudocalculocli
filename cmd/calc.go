/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"calc/app"
	"fmt"
	"log"
	"strconv"

	"github.com/spf13/cobra"
)

// calcCmd represents the calc command

var calcCmd = &cobra.Command{
	Use:   "calc",
	Short: "A brief description of your command",
	Long:  `retornar o calculo`,
	Run: func(cmd *cobra.Command, args []string) {
		log.Println(">>>", args)

		a, _ := strconv.ParseFloat(args[0], 64)

		b, _ := strconv.ParseFloat(args[1], 64)

		calc := app.NewCalc()
		calc.A = a
		calc.B = b
		log.Println(">>>", calc)

		fmt.Println("o valor é", calc.Sum())

	},
}

////func init() {
//rootCmd.AddCommand(calcCmd)

// Here you will define your flags and configuration settings.

// Cobra supports Persistent Flags which will work for this command
// and all subcommands, e.g.:
// calcCmd.PersistentFlags().String("foo", "", "A help for foo")

// Cobra supports local flags which will only run when this command
// is called directly, e.g.:
// calcCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
//}
