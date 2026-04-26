package main

import "fmt"

func main() {
	survay := &Survay{
		Title:     "Customer Satisfaction",
		questions: []string{"How satisfied are you with our service?", "Would you recommend us to others?"},
	}
	repo := &InMemoryRepository{}
	err := SaveSurvay(survay, repo)
	if err != nil {
		panic(err)
	}
	fmt.Println(repo.GetAll()[0].getTitle())
}

// type survey serves single responsibility of holding survey data and providing access to it,
// while the repository interface and its implementation handle the responsibility of saving and retrieving surveys.
//  This separation of concerns adheres to the Single Responsibility Principle (SRP).
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

type Repository interface {
	Save(s *Survay) error
}

type InMemoryRepository struct {
	survays []*Survay
}

func (r *InMemoryRepository) Save(s *Survay) error {
	r.survays = append(r.survays, s)
	return nil
}

func (r *InMemoryRepository) GetAll() []*Survay {
	return r.survays
}

func SaveSurvay(s *Survay, r Repository) error {
	return r.Save(s)
}
