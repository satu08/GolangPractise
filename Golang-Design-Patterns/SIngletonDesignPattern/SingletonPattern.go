// Singleton Design Pattern in Go?
// 1. What ?
// singleton design pattern is creational design pattern that ensures a class has only one instance and provides a global point of access to it.
//  It is used when we want to restrict the instantiation of a class to a single object and provide a global access point to that object.
//  This can be useful in situations where we want to control access to shared resources, such as database connections or configuration settings.
// db instance
// config instance
// logger instance
// cache instance

// 2. Why ?

// Singleton ensures a single instance of a type across the application.
// In Go, it is implemented using sync.Once for thread-safe lazy initialization.
//  However, Go developers often prefer dependency injection over singletons to avoid global state and improve testability.
// 3. How ?

// | Feature               | Singleton ❌ | Dependency Injection ✅ |
// | --------------------- | ----------- | ---------------------- |
// | Dependency visibility | Hidden      | Explicit               |
// | Testability           | Hard        | Easy                   |
// | Flexibility           | Low         | High                   |
// | Coupling              | Tight       | Loose                  |
// | Scalability           | Poor        | Excellent              |

package main

import (
	"fmt"
	"sync"
)

var mutex = &sync.Mutex{}

type config struct {
	// Config variables
}

var counter int = 1
var singleConfigInstance *config

func getConfigInstance() *config {
	if singleConfigInstance == nil {
		mutex.Lock()
		defer mutex.Unlock()
		if singleConfigInstance == nil {
			fmt.Println("Creting Single Instance Now, and Counter is ", counter)
			singleConfigInstance = &config{}
			counter = counter + 1
		} else {
			fmt.Println("Single Instance already created-1, returning that one")
		}
	} else {
		fmt.Println("Single Instance already created-2, returning the same")
	}
	return singleConfigInstance
}

func main() {
	// for i := 0; i < 100; i++ {
	// 	go getConfigInstance()
	// }

	// d1 := GetDB()
	// d2 := GetDB()
	// fmt.Println(d1 == d2)
	// fmt.Scanln()
	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			GetDB()
		}()
	}
	wg.Wait()
}

type DB struct{}

var dbInstance *DB
var once sync.Once

func GetDB() *DB {
	once.Do(func() {
		fmt.Println("Creating DB instance")
		dbInstance = &DB{}
	})
	fmt.Println("Returning DB instance")
	return dbInstance
}

type ShapeI interface {
	Shape() string
}

type Circle struct{}

func (c Circle) Shape() string {
	return "Circle"
}

type Square struct{}

func (s Square) Shape() string {
	return "Square"
}

func newCircle() ShapeI {
	return &Circle{}
}

func newSquare() ShapeI {
	return &Square{}
}
