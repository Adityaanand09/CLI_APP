/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// ccwcCmd represents the ccwc command
var ccwcCmd = &cobra.Command{
	Use:   "ccwc",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("ccwc called")
		flagValue, name := cmd.Flags().GetString("c")
		fmt.Println(flagValue, " ", name)
		if flagValue != "" {
			fileInfo, err := os.Stat(flagValue)
			if err != nil {
				fmt.Println("Error in Getting File, Please check the path")
			}
			fmt.Println(fileInfo.Size())
		}
	},
}

func init() {
	rootCmd.AddCommand(ccwcCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	ccwcCmd.PersistentFlags().String("foo", "", "A help for foo")
	ccwcCmd.PersistentFlags().String("c", "", "Gives the Number of Bytes in a File")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// ccwcCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
