package main

import "fmt"

// Decorator pattern allows adding behavior to an object dynamically by wrapping it with another object that implements the same interface.
//  In Go, it is commonly implemented using function wrappers or interface-based composition, and is widely used in HTTP middleware.

type Notifier interface {
	Send(msg string)
}

type Email struct{}

func (e Email) Send(msg string) {
	fmt.Println("Email:", msg)
}

// Decorator
type LoggingDecorator struct {
	wrapped Notifier
}

func (l LoggingDecorator) Send(msg string) {
	fmt.Println("Logging:", msg)
	l.wrapped.Send(msg)
}

func main() {
	email := Email{}
	loggingEmail := LoggingDecorator{wrapped: email}
	loggingEmail.Send("Hello, World!")
}
