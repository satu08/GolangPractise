package main

import (
	"encoding/json"
	"io"
	"os"
)

func main() {
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

func SendSurvay(s *Survay, w io.Writer) error {
	b, err := json.Marshal(s)
	if err != nil {
		return err
	}
	_, err = w.Write(b)
	return err
}

// The Writer interface defines a single method Write, which takes a byte slice and returns the number of bytes written and an error.
// The FileWriter struct implements the Writer interface, providing a concrete implementation of the Write method.
// The SendSurvay function takes a Survay and a Writer, and calls the Write method on the Writer, allowing for flexible output of survey data to various destinations (e.g., files, standard output, network) without modifying existing code,
// adhering to the liskov Substitution Principle (LSP) and Interface Segregation Principle (ISP).
type Writer interface {
	Write(p []byte) (n int, err error)
}

type FileWriter struct {
}

func (f *FileWriter) Write(p []byte) (n int, err error) {
	return 0, nil

}

func example() {
	file, err := os.Create("survay.json")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	survay := &Survay{
		Title:     "Customer Satisfaction",
		questions: []string{"How satisfied are you with our service?", "Would you recommend us to others?"},
	}
	err = SendSurvay(survay, file)
	if err != nil {
		panic(err)
	}

	err = SendSurvay(survay, os.Stdout)
	if err != nil {
		panic(err)
	}

	err = SendSurvay(survay, &FileWriter{})
	if err != nil {
		panic(err)
	}
}
