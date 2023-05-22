/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"os"

	"github.com/spf13/cobra"
)

// readDirCmd represents the readDir command
var readDirCmd = &cobra.Command{
	Use:   "readDir",
	Short: "list directory contents",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		c, err := os.ReadDir(args[0])
		verifyError(err)
		for _, dir := range c {
			fmt.Println(" ", dir.Name(), dir.IsDir())
		}
	},
}

func init() {
	rootCmd.AddCommand(readDirCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// readDirCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// readDirCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
