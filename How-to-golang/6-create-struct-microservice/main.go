package main

import (
	"fmt"
)

func main() {
	fmt.Println("creating and structuring microservice")
	svc := NewCatFactService("https://catfact.ninja/fact")
	svc = NewLoggingService(svc)
	apiserv := NewApiServer(svc)
	apiserv.Start(":3001")
}
