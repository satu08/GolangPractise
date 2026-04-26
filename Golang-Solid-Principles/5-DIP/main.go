package main

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
	//fmt.Println(repo.GetAll()[0].getTitle())
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

// func (r *InMemoryRepository) GetAll() []*Survay {
// 	return r.survays
// }

func SaveSurvay(s *Survay, r Repository) error {
	return r.Save(s)
}
// The SurveyManager struct serves as a higher-level abstraction that manages surveys using a Repository, 
// adhering to the Dependency Inversion Principle (DIP) by depending on abstractions rather than concrete implementations.

type SurveyManager struct {
	store Repository
}

func newSurveyManager(stor Repository) *SurveyManager {
	return &SurveyManager{
		store: stor,
	}
}
