/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var git bool

// createCmd represents the create command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Creates a new list for this branch or folder",
	RunE: func(cmd *cobra.Command, args []string) error {
		err := Instance.CreateList(cmd.Context(), git)
		if err != nil {
			return fmt.Errorf("could not create list: %s", err)
		}
		if git {
			fmt.Println("Created a list for this branch")
			return nil
		}
		fmt.Println("Created a list for this folder")

		return nil
	},
}

func init() {
	rootCmd.AddCommand(initCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// createCmd.PersistentFlags().String("foo", "", "A help for foo")
	listCmd.Flags().BoolVarP(&git, "no-git", "g", true, "Use to disable linking with git branches")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// createCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
