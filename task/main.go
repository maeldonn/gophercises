package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/maeldonn/gophercises/task/cmd"
	"github.com/maeldonn/gophercises/task/db"
)

func main() {
	home, _ := os.UserHomeDir()
	dbPath := filepath.Join(home, "task.db")
	must(db.Init(dbPath))
	must(cmd.RootCmd.Execute())
}

func must(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
