package cmd

import (
	"fmt"
	"os"

	"github.com/maeldonn/gophercises/task/db"
	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(completedCmd)
}

var completedCmd = &cobra.Command{
	Use:   "completed",
	Short: "Lists of all your completed tasks",
	Run: func(cmd *cobra.Command, args []string) {
		tasks, err := db.CompletedTasks()
		if err != nil {
			fmt.Println("Something went wrong:", err)
			os.Exit(1)
		}

		if len(tasks) == 0 {
			fmt.Println("You have no tasks completed! Go to work!")
			return
		}

		fmt.Println("You have completed the following tasks:")
		for i, task := range tasks {
			fmt.Printf("%d. %s (DONE)\n", i+1, task.Value)
		}
	},
}
