package main

import (
	"fmt"
	"log"
	"math"
)

type norgatemath struct {
	err error
}

func (n norgatemath) Error() string {
	return fmt.Sprintf("a norgate math error: %v", n.err)
}

func main() {
	fmt.Println("info for errors")

	v, err := sqrt2(-10)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(v)
}

func sqrt2(f float64) (float64, error) {
	if f < 0 {
		nme := fmt.Errorf("-- cannot sqrt: negative number : %v ", f)
		return 0, norgatemath{nme}
	}
	return math.Sqrt(f), nil
}
