package quiz

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

type Quiz struct {
	problems []Problem
	score    int
	limit    int
	answer   chan string
}

func New(p []Problem, l int) *Quiz {
	return &Quiz{
		problems: p,
		limit:    l,
		answer:   make(chan string, 1),
	}
}

func (q *Quiz) Start() {
	timer := time.NewTimer(time.Second * time.Duration(q.limit))
	for i, p := range q.problems {
		fmt.Printf("Problem #%d: %s = \n", i, p.question)

		go q.scan()

		select {
		case <-timer.C:
			return
		case ans := <-q.answer:
			if strings.ToLower(p.answer) == strings.ToLower(ans) {
				q.score++
			}
		}
	}
}

func (q *Quiz) scan() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	q.answer <- scanner.Text()
}

func (q *Quiz) DisplayScore() {
	fmt.Printf("You scored %d out of %d", q.score, len(q.problems))
}
