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
		flagValue, _ := cmd.Flags().GetString("s")
		//fmt.Println(flagValue, " ", name)
		if flagValue != "" {
			fileInfo, err := os.Stat(flagValue)
			if err != nil {
				fmt.Println("Error in Getting File, Please check the path")
			}
			fmt.Println(fileInfo.Size())
		}
	},
}

var countLineCmd = &cobra.Command{
	Use:   "cl",
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

var countwords = &cobra.Command{
	Use:   "wc",
	Short: "Count words in a file",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		flagValue, _ := cmd.Flags().GetString("w")
		if flagValue != "" {
			reader, _ := os.ReadFile(flagValue)
			content := string(reader)
			wordCount := 0
			contentWithSpace := content + " "
			for _, a := range contentWithSpace {
				if a == ' ' {
					wordCount++
				}
			}
			fmt.Println(wordCount)
		}
	},
}

func init() {
	rootCmd.AddCommand(ccwcCmd)
	rootCmd.AddCommand(countLineCmd)
	rootCmd.AddCommand(countwords)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	ccwcCmd.PersistentFlags().String("foo", "", "A help for foo")
	ccwcCmd.PersistentFlags().String("s", "", "Gives the Number of Bytes in a File")
	countLineCmd.PersistentFlags().String("l", "", "Counts the Number of Lines in a file")
	countwords.PersistentFlags().String("w", "", "Count words in a file")
}
