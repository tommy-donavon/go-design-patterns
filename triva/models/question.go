package models

import (
	"fmt"
	"math/rand"
	"triva/interfaces"
	"triva/utils"

	"github.com/fatih/color"
)

var (
	Score = 0
	OutOf = 0
)

type Question struct {
	question    string
	correctAns  string
	possibleAns []string
}

func NewQuestion(question string, correctAnswer string, otherAnswers ...string) *Question {
	otherAnswers = append(otherAnswers, correctAnswer)
	for i := range otherAnswers {
		j := rand.Intn(i + 1)
		otherAnswers[i], otherAnswers[j] = otherAnswers[j], otherAnswers[i]
	}
	return &Question{
		question:    question,
		correctAns:  correctAnswer,
		possibleAns: otherAnswers,
	}
}

func (q *Question) AskQuestion() error {
	fmt.Println(q.question)
	userChoice := utils.MakeMenu(q.possibleAns...)
	OutOf += 5
	if userChoice == 0 {
		return fmt.Errorf("user exited")
	}
	answer := q.possibleAns[userChoice-1]
	if answer == q.correctAns {
		c := color.New(color.FgGreen)
		c.Println("You got the answer right!!")
		Score += 5
	} else {
		c := color.New(color.FgRed)
		c.Printf("Wrong! the right answer was %s\n", q.correctAns)
	}

	return nil
}

func (q *Question) GetIterator() (interfaces.IIterator, error) {
	return &leafIterator{}, nil
}
func (q *Question) Add(c interfaces.IComponent) error {
	return fmt.Errorf("i am a leaf")
}
func (q *Question) Remove(c interfaces.IComponent) error {
	return fmt.Errorf("i am a leaf")
}
