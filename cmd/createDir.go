/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// createDirCmd represents the createDir command
var createDirCmd = &cobra.Command{
	Use:   "createDir",
	Short: "Create a new directory",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		err := os.Chdir("../..")
		verifyError(err)
		err = os.MkdirAll(args[0], 0777)
		verifyError(err)
	},
}

func init() {
	rootCmd.AddCommand(createDirCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// createDirCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// createDirCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
