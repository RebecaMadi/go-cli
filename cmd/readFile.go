/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"os"

	"github.com/spf13/cobra"
)

// readFileCmd represents the readFile command
var readFileCmd = &cobra.Command{
	Use:   "readFile",
	Short: "Read a file",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, str []string) {
		dr, err := os.ReadFile(string(str[0]))
		verifyError(err)
		fmt.Println(string(dr))
	},
}

func verifyError(e error) {
	if e != nil {
		fmt.Println(e.Error())
	}
}

func init() {
	rootCmd.AddCommand(readFileCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// readFileCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// readFileCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
