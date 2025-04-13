package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/Beadko/gophercises_task_manager/db"
	"github.com/spf13/cobra"
)

// addCmd represent the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Adds a task to your task list",
	Run: func(cmd *cobra.Command, args []string) {
		t := strings.Join(args, " ")
		k, err := db.CreateTask(t)
		if err != nil {
			fmt.Println("Something went wrong:", err)
			os.Exit(1)
		}
		fmt.Printf("Added \"%d. %s\" to your task list.\n", k, t)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
