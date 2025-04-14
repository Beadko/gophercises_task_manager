package cmd

import (
	"fmt"
	"os"

	"github.com/Beadko/gophercises_task_manager/db"
	"github.com/spf13/cobra"
)

// complCmd represents the list of the completed tasks command
var complCmd = &cobra.Command{
	Use:   "completed",
	Short: "List all of the tasks completed today",
	Run: func(cmd *cobra.Command, args []string) {
		tasks, err := db.CompletedTasks()
		if err != nil {
			fmt.Println("Something went wrong:", err)
			os.Exit(1)
		}
		if len(tasks) == 0 {
			fmt.Println("You have not completed any tasks today")
			return
		}
		fmt.Println("You have finished these tasks today:")
		for i, t := range tasks {
			fmt.Printf("%d. %s\n", i+1, t.Value)
		}
	},
}

func init() {
	rootCmd.AddCommand(complCmd)
}
