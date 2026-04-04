package main

import "fmt"

func main() {
	prime := isPrime(9)
	if prime == true {
		fmt.Println("its a prime number")
	} else {
		fmt.Println("its not a prime number")
	}
	primeFactors(12)
}

func isPrime(n int) bool {
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func primeFactors(n int) {
	for i := 2; i < n; i++ {
		if isPrime(i) {
			x := i
			for n%x == 0 {
				fmt.Println(i)
				x *= i
			}
		}
	}
}
