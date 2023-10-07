/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// countLineCmd represents the countLine command
var countLineCmd = &cobra.Command{
	Use:   "countLine",
	Short: "Counts the number of lines in a file",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("countLine called")
		flagValue, _ := cmd.Flags().GetString("l")
		if flagValue != "" {
			file, err := os.Stat(flagValue)
			if err != nil {
				fmt.Println("Error in Getting File, Please check the path")
			}
			reader, _ := os.ReadFile(flagValue)
			content := string(reader)
			fmt.Println(content)
			lineCount := 0
			fmt.Println(file, " ", flagValue)

			for _, a := range content {
				if a == '\n' {
					fmt.Println(a)
					lineCount++
				}
			}
			fmt.Println(lineCount)
		}

	},
}

func init() {
	rootCmd.AddCommand(countLineCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	countLineCmd.PersistentFlags().String("l", "", "Counts the Number of Lines in a file")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// countLineCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
