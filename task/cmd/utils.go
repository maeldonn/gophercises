package cmd

import (
	"fmt"
	"strconv"
)

func parseIds(args []string) []int {
	var ids []int
	for _, arg := range args {
		id, err := strconv.Atoi(arg)
		if err != nil {
			fmt.Println("Failed to parse the argument:", arg)
		} else {
			ids = append(ids, id)
		}
	}
	return ids
}
