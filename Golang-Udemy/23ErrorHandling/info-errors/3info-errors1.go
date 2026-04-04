package main

import (
	"errors"
	"fmt"
	"log"
	"math"
)

func main() {
	fmt.Println("info for errors")

	_, err := sqrt(-10)
	if err != nil {
		log.Fatalln(err)
	}
}

func sqrt(f float64) (float64, error) {
	if f < 0 {
		return 0, errors.New("cannot sqrt: negative number")
	}
	return math.Sqrt(f), nil
}
