// Chain of Responsibility Design Pattern

// WHAT
// 👉 Chain of Responsibility = Pass a request through a chain of handlers until one handles it
//Chain of Responsibility pattern allows a request to pass through a sequence of handlers,
//  where each handler decides whether to process the request or pass it to the next.
//  In Go, it is commonly implemented using interfaces and struct composition and is widely used in HTTP middleware pipelines.
// WHY
// HOW

package main

import "fmt"

type step interface {
	execute()     // Func to be executed next.
	setNext(step) // Func to set the next func to be executed.
}

type openFlap struct {
	next step
}

type pressStartButton struct {
	next step
}

type enterPassword struct {
	next step
}

// Open Flap

func (r *openFlap) execute() {
	fmt.Println("Opening Laptop Flap")
	r.next.execute()
}

func (r *openFlap) setNext(next step) {
	r.next = next
}

// Press Button Below

func (d *pressStartButton) execute() {
	fmt.Println("Pressing start button")
	d.next.execute()
}

func (d *pressStartButton) setNext(next step) {
	d.next = next
}

// Enter Password

func (m *enterPassword) execute() {
	fmt.Println("Entering password for Laptop")
	fmt.Println("Laptop Opened!")
}

func (m *enterPassword) setNext(next step) {
	m.next = next
}

func main() {
	//Set next for enterPassword step
	enterPassword := &enterPassword{}

	//Set next for pressStartButton step
	pressStartButton := &pressStartButton{}
	pressStartButton.setNext(enterPassword)

	//Set next for openFlap step
	openFlap := &openFlap{}
	openFlap.setNext(pressStartButton)

	openFlap.execute()

	auth := &AuthHandler{}
	log := &LogHandler{}

	auth.SetNext(log)

	auth.Handle("log")
}

//////////////////// HANDLER INTERFACE ////////////////////

type Handler interface {
	SetNext(Handler)
	Handle(string)
}

//////////////////// BASE HANDLER ////////////////////

type BaseHandler struct {
	next Handler
}

func (b *BaseHandler) SetNext(n Handler) {
	b.next = n
}

//////////////////// AUTH HANDLER ////////////////////

type AuthHandler struct {
	BaseHandler
}

func (a AuthHandler) Handle(req string) {
	if req == "auth" {
		fmt.Println("Auth handled ✅")
		return
	}
	if a.next != nil {
		a.next.Handle(req)
	}
}

//////////////////// LOG HANDLER ////////////////////

type LogHandler struct {
	BaseHandler
}

func (l LogHandler) Handle(req string) {
	if req == "log" {
		fmt.Println("Log handled ✅")
		return
	}
	if l.next != nil {
		l.next.Handle(req)
	}
}
