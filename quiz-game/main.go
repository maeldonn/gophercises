package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/maeldonn/gophercises/quiz-game/quiz"
)

var (
	csvFilname string
	limit      int
)

func init() {
	flag.StringVar(&csvFilname, "csv", "problems.csv", "a csv file in a format of 'question, answer'")
	flag.IntVar(&limit, "limit", 30, "The time limit for the quiz in seconds")
	flag.Parse()
}

func main() {
	var (
		lines    = readCsv(csvFilname)
		problems = parseLines(lines)
	)

	quiz := quiz.New(problems, limit)
	quiz.Start()
	quiz.DisplayScore()
}

func readCsv(filename string) [][]string {
	file, err := os.Open(filename)
	if err != nil {
		exit(fmt.Sprintf("Failed to open the CSV file: %s", filename))
	}

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		exit("Failed to parse the provided CSV file")
	}

	return records
}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}

func parseLines(lines [][]string) []quiz.Problem {
	p := make([]quiz.Problem, len(lines))
	for i, line := range lines {
		p[i] = quiz.NewProblem(line[0], strings.TrimSpace(line[1]))
	}
	return p
}
