package cmd

import (
	"fmt"
	"os"
	"strconv"

	"github.com/Beadko/gophercises_task_manager/db"
	"github.com/spf13/cobra"
)

// rmCmd represents the remove task command
var rmCmd = &cobra.Command{
	Use:   "rm",
	Short: "Deletes the task",
	Run: func(cmd *cobra.Command, args []string) {
		var ids []int
		for _, arg := range args {
			id, err := strconv.Atoi(arg)
			if err != nil {
				fmt.Println("Failed to parse the argument:", arg)
			} else {
				ids = append(ids, id)
			}
		}
		tasks, err := db.AllTasks()
		if err != nil {
			fmt.Println("Something went wrong:", err)
			os.Exit(1)
		}
		for _, id := range ids {
			if id <= 0 || id > len(tasks) {
				fmt.Println("Invalid task number:", id)
				continue
			}
			task := tasks[id-1]
			err := db.DeleteTask(task.Key)
			if err != nil {
				fmt.Printf("Failed to delete \"%d\". Error :%s\n", id, err)
			} else {
				fmt.Printf("Deleted \"%d\".\n", id)

			}
		}
	},
}

func init() {
	rootCmd.AddCommand(rmCmd)
}
