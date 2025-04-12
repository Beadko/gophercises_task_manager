package cmd

import (
	"fmt"
	"os"

	"github.com/Beadko/gophercises_task_manager/db"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all of your tasks",
	Run: func(cmd *cobra.Command, args []string) {
		tasks, err := db.AllTasks()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Something went wrong:%v\n", err)
			os.Exit(1)
		}
		if len(tasks) == 0 {
			fmt.Println("You have completed all tasks. Well done! 😎")
			return
		}
		fmt.Println("You have the following tasks:")
		for i, t := range tasks {
			fmt.Printf("%d. %s\n", i+1, t.Value)
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
