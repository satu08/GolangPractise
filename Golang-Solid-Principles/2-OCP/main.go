package main

func main() {
	s := &Survay{
		Title:     "Customer Satisfaction",
		questions: []string{"How satisfied are you with our service?", "Would you recommend us to others?"},
	}
	e := &s3{}
	err := ExportSurvay(s, e)
	if err != nil {
		panic(err)
	}
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

// The Exporter interface defines a single method Export, which takes a Survay object and returns an error.
// The s3, csv, and pdf types implement the Exporter interface, each providing its own logic for exporting the survey data.
// The ExportSurvay function takes a Survay and an Exporter, and calls the Export method on the Exporter, allowing for flexible exporting of survey data without modifying existing code,
// adhering to the Open/Closed Principle (OCP).
type Exporter interface {
	Export(s *Survay) error
}

type s3 struct {
}

func (s *s3) Export(surv *Survay) error { // logic to export survey to s3
	return nil
}

type csv struct {
}

func (c *csv) Export(surv *Survay) error {
	// logic to export survey to csv
	return nil
}

type pdf struct {
}

func (p *pdf) Export(surv *Survay) error {
	// logic to export survey to pdf
	return nil
}

func ExportSurvay(s *Survay, e Exporter) error {
	return e.Export(s)
}
