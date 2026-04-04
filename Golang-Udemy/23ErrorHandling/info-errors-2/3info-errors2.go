package main

import (
	"fmt"
	"log"
	"math"
)

func main() {
	fmt.Println("info for errors")

	v, err := sqrt1(-10)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(v)
}

func sqrt1(f float64) (float64, error) {
	if f < 0 {
		return 0, fmt.Errorf("cannot sqrt: negative number : %v ", f)
	}
	return math.Sqrt(f), nil
}
