package cmd

import (
	"fmt"
	"os"

	"github.com/maeldonn/gophercises/task/db"
	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(rmCmd)
}

var rmCmd = &cobra.Command{
	Use:   "rm",
	Short: "Delete a task from your task list",
	Run: func(cmd *cobra.Command, args []string) {
		ids := parseIds(args)
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
				fmt.Printf("Failed to delete \"%d\". Error: %s\n", id, err)
			} else {
				fmt.Printf("Deleted \"%d\"\n", id)
			}
		}
	},
}
