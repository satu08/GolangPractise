package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type Person struct {
	First string
	Last  string
	Age   int
}

func main() {
	fmt.Println("introduction to json marshal")
	p1 := Person{
		First: "satya",
		Last:  "jadhav",
		Age:   23,
	}
	p2 := Person{
		First: "saurabh",
		Last:  "vidhate",
		Age:   24,
	}
	people := []Person{p1, p2}
	fmt.Println(people)
	bs, err := json.Marshal(people)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(bs))
	f, err := os.Create("output.json")
	if err != nil {
		log.Fatalf("%s", err)
	}
	defer f.Close()

	fmt.Println("-------------------------")
	err = json.NewEncoder(f).Encode(people)
	if err != nil {
		fmt.Println(err)
	}
}
