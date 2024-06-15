package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

func main() {

	n, err := fmt.Println("hello world")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(n) // returns number of bytes

	var ans1, ans2, ans3 string

	fmt.Println("enter name:")
	_, err = fmt.Scan(&ans1)
	if err != nil {
		panic(err)
	}
	fmt.Println("enter favorite car:")
	_, err = fmt.Scan(&ans2)
	if err != nil {
		panic(err)
	}
	fmt.Println("Enter favorite color")
	_, err = fmt.Scan(&ans3)
	if err != nil {
		panic(err)
	}

	fmt.Println(ans1, ans2, ans3)

	f, err1 := os.Create("output.txt")
	if err1 != nil {
		fmt.Println(err1)
		return
	}
	defer f.Close()

	r := strings.NewReader(ans1)
	_, err2 := io.Copy(f, r)
	if err2 != nil {
		fmt.Println(err2)
	}

	f, err = os.Open("output.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()

	bs, err3 := ioutil.ReadAll(f)
	if err3 != nil {
		fmt.Println(err3)
	}
	fmt.Println(string(bs))

}
