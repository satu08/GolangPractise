package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
)

type person struct {
	firstname string
}

func (p person) writeout(w io.Writer) error {
	_, err := w.Write([]byte(p.firstname))
	return err
}
func main() {
	p := person{
		firstname: "jadhav",
	}
	f, err := os.Create("output.txt")
	if err != nil {
		log.Fatalf("%s", err)
	}
	defer f.Close()
	var b bytes.Buffer
	err = p.writeout(&b)
	if err != nil {
		return
	}
	err = p.writeout(f)
	if err != nil {
		return
	}
	fmt.Println(b.String())

	//s := []byte("satya")
	//
	//_, err = f.Write(s)
	//if err != nil {
	//	log.Fatalf("%s", err)
	//}

	// using buffer package
	z := bytes.NewBufferString("satya")
	fmt.Println(z.String())
	z.WriteString("jadhav")
	fmt.Println(z.String())
	z.Reset()
	z.WriteString("ravi")
	fmt.Println(z.String())
	z.Write([]byte("satya"))
	fmt.Println(z.String())

}
