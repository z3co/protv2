/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

// completeCmd represents the complete command
var completeCmd = &cobra.Command{
	Use:   "complete [id]",
	Short: "Changes the status of the todo with the id passed",
	Args:  cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		id, err := strconv.Atoi(args[0])
		if err != nil {
			return fmt.Errorf("wrong input it is not an int")
		}
		err = Instance.UpdateStatus(cmd.Context(), int64(id))
		if err != nil {
			return fmt.Errorf("could not change status: %s", err)
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(completeCmd)
}
