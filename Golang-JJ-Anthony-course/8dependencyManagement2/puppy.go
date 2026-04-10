package dependencymanagement2

import dependencymanagement3 "hello/Golang-JJ-Anthony-course/8dependencymanagement3"

func Bark() string {
	return "woof!"
}

func Barks() string {
	return "woof! woof! woof!"
}

func BigBarks() string {
	return dependencymanagement3.DogBarks(Bark())
}
