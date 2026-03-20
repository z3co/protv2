/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/spf13/cobra"
	db "github.com/z3co/prot/db/gen"
)

var all bool

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all todos for this branch/folder",
	RunE: func(cmd *cobra.Command, args []string) error {
		allTodos, err := Instance.GetTodos(cmd.Context())
		if err != nil {
			return err
		}
		var todos []db.Todo
		if !all {
			for _, row := range allTodos {
				if row.Done == 0 {
					todos = append(todos, row)
				}
			}
		} else {
			todos = allTodos
		}
		w := tabwriter.NewWriter(os.Stdout, 0, 0, 4, ' ', 0)
		fmt.Fprintln(w, "Id\tDescription\tDone")
		for _, row := range todos {
			fmt.Fprintf(w, "%v\t%s\t%v\n", row.ID, row.Description, row.Done == 1)
		}
		w.Flush()
		return nil
	},
}

func init() {
	rootCmd.AddCommand(listCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	listCmd.Flags().BoolVarP(&all, "all", "a", false, "Displays all todos instead of completed ones")
}
