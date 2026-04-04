package main

import (
	"os"
)

func main() {
	_, err := os.Open("test1.log")
	if err != nil {
		//fmt.Println("Error occurred - ", err)
		//log.Println("Error occurred - ", err)
		//log.Fatalln(err)  // program shut down immediately and no defer function will run
		//log.Panicln(err) // exit from current goroutine
		panic(err)
	}

	//f, err1 := os.Create("test.log")
	//if err1 != nil {
	//	log.Println("Error occurred - ", err1)
	//}
	//defer f.Close()
	//log.SetOutput(f)
	//
	//f2, err2 := os.Open("test1.log")
	//if err2 != nil {
	//	log.Println("Error occurred - ", err2)
	//}
	//defer f2.Close()
	//fmt.Println("see the output in test.log file")

}
