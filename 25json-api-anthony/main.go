package main

import (
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func main() {
	store, err := NewPostgresStore()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("satyanarayan jadhav")
	fmt.Printf("%+v\n", store)
	// server := NewAPIServer(":3001", store)
	// server.Run()
}
