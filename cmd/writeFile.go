/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"
	"github.com/spf13/cobra"
)

// writeFileCmd represents the writeFile command
var writeFileCmd = &cobra.Command{
	Use:   "writeFile",
	Short: "Update a file",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		//err := os.Chdir("../..")
		file, err := os.OpenFile(args[0], os.O_APPEND|os.O_WRONLY, 0777)
		verifyError(err)
		defer file.Close()
		for _, v := range args[1:]{
			_, err := file.WriteString(v + " ")
			verifyError(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(writeFileCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// writeFileCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// writeFileCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
