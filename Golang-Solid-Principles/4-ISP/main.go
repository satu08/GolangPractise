package main

import "fmt"

func main() {
	fmt.Println("Interface Segregation Principle (ISP) Example")

}

type Survay struct {
	Title     string
	questions []string
}

func (s *Survay) getTitle() string {
	return s.Title
}

func (s *Survay) getQuestions() []string {
	return s.questions
}

type Question interface {
	SetTitle(title string)
}

type QuestionWithOptions interface {
	Question
	SetOptions(options []string)
}

type textQuestion struct {
	Title string
}

func (t *textQuestion) SetTitle(title string) {
	t.Title = title
}

type choiceQuestion struct {
	Title   string
	Options []string
}

func (c *choiceQuestion) SetTitle(title string) {
	c.Title = title
}

func (c *choiceQuestion) SetOptions(options []string) {
	c.Options = options
}

// The Question interface defines a method for setting the title of a question, while the QuestionWithOptions interface extends it to include a method for setting options.
// The textQuestion struct implements the Question interface, while the choiceQuestion struct implements both interfaces, allowing for different types of questions to be created without forcing all question types to implement methods they don't need,
// adhering to the Interface Segregation Principle (ISP).
func example() {
	textQ := &textQuestion{}
	textQ.SetTitle("What is your name?")
	choiceQ := &choiceQuestion{}
	choiceQ.SetTitle("What is your favorite color?")
	choiceQ.SetOptions([]string{"Red", "Green", "Blue"})
}
