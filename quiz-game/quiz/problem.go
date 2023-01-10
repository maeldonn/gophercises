package quiz

type Problem struct {
	question string
	answer   string
}

func NewProblem(question, answer string) Problem {
	return Problem{
		question: question,
		answer:   answer,
	}
}
