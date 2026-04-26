// Factory Design Pattern in Go:
// Why
// What
// Factory design pattern is a creational design pattern that provides an interface for creating objects in a superclass,
//  but allows subclasses to alter the type of objects that will be created.
// It is used when we want to create objects without specifying the exact class of object that will be created.
//  This can be useful in situations where we want to decouple the client code from the implementation details of the objects being created,
// or when we want to provide a common interface for creating different types of objects.
//
// How

// Factory pattern encapsulates object creation logic and returns an interface instead of concrete types,
//  promoting loose coupling and flexibility.
// In Go, it is commonly implemented using constructor functions like NewX() and often combined with dependency injection.

package main

import "fmt"

// App
type Car interface {
	getCar() string
}

type SedanType struct {
	Name string
}

func getNewSedan() *SedanType {
	return &SedanType{}
}

func (s *SedanType) getCar() string {
	return "Honda City"
}

type HatchBackType struct {
	Name string
}

func getNewHatchBack() *HatchBackType {
	return &HatchBackType{}
}

func (h *HatchBackType) getCar() string {
	return "Polo"
}

func main() { // Client
	getCarFactory(1)
	getCarFactory(2)

	db := NewDatabase("mysql") // client doesn't care about struct
	db.Connect()

	db = NewDatabase("postgres")
	db.Connect()
}

func getCarFactory(cartype int) { // Factory Class/Object
	var car Car
	if cartype == 1 {
		car = getNewHatchBack()
	} else {
		car = getNewSedan()
	}
	carInfo := car.getCar()
	fmt.Println(carInfo)

}

//////////////////// INTERFACE ////////////////////

type Database interface {
	Connect()
}

//////////////////// IMPLEMENTATIONS ////////////////////

type MySQL struct{}

func (m MySQL) Connect() {
	fmt.Println("Connected to MySQL")
}

type Postgres struct{}

func (p Postgres) Connect() {
	fmt.Println("Connected to Postgres")
}

//////////////////// FACTORY ////////////////////

func NewDatabase(dbType string) Database {
	switch dbType {
	case "mysql":
		return MySQL{}
	case "postgres":
		return Postgres{}
	default:
		return nil
	}
}
