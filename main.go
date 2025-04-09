package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/Beadko/gophercises_task_manager/cmd"
	"github.com/Beadko/gophercises_task_manager/db"
)

func main() {
	home, err := os.UserHomeDir()
	if err != nil {
		fmt.Printf("Could not get the user home directory: %v\n", err)
		return
	}
	dbPath := filepath.Join(home, "tasks.db")
	if err := db.Init(dbPath); err != nil {
		fmt.Printf("Could not initialise the database: %v\n", err)
		return
	}
	cmd.Execute()
}
