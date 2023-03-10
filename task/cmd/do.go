package cmd

import (
	"fmt"
	"os"

	"github.com/maeldonn/gophercises/task/db"
	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(doCmd)
}

var doCmd = &cobra.Command{
	Use:   "do",
	Short: "Mark a task as completed",
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
			err := db.CompleteTask(task.Key)
			if err != nil {
				fmt.Printf("Failed to complete \"%d\". Error: %s\n", id, err)
			} else {
				fmt.Printf("Completed \"%d\"\n", id)
			}
		}
	},
}
