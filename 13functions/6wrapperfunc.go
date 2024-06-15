package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

func main() {

	bytes, err := readFile("../15writerInterfaces/output.txt")
	if err != nil {
		log.Fatalf("%s", err)
	}
	fmt.Println(bytes)
	fmt.Println(string(bytes))

	timeFunc(doWork)

}

func readFile(filename string) ([]byte, error) {
	xb, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("error in readfile func : %s", err)
	}
	return xb, nil
}

func doWork() {
	for i := 0; i < 2000; i++ {
		fmt.Println(i)
	}
}

func timeFunc(f func()) {
	start := time.Now()
	f()
	elapsed := time.Since(start)
	fmt.Println(elapsed)
}
