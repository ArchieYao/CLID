/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

// sumCmd represents the sum command
var sumCmd = &cobra.Command{
	Use:   "sum",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("sum called")
		result := 0
		for _, v := range args {
			arg, err := strconv.Atoi(v)
			if err != nil {
				fmt.Printf("args: %s , %s cannot convert to int ", args, v)
				return
			}
			result += arg
		}

		fmt.Printf("result [%d]", result)
	},
}

func init() {
	rootCmd.AddCommand(sumCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// sumCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// sumCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
