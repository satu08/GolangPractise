package main

import "fmt"

func main() {
	lcm := lcm(4, 6)
	fmt.Println(lcm)
	gcd := gcd(4, 6)
	fmt.Println(gcd)
	fmt.Println(lcm2(4, 6))
	fmt.Println(euclideanAlgo(6, 12))
	fmt.Println(optimisedEuclideanAlgo(6, 12))
}

func lcm(a, b int) int {
	return a * b / gcd(a, b)
}

func gcd(a, b int) int {
	min := min(a, b)
	for min > 0 {
		if a%min == 0 && b%min == 0 {
			return min
		}
		min--
	}
	return 1
}

func lcm2(a, b int) int {
	maxi := max(a, b)
	for {
		if maxi%a == 0 && maxi%b == 0 {
			return maxi
		}
		maxi++

	}
}

// gcd(a,b)=gcd(a−b,b)
func euclideanAlgo(a, b int) int {
	for a != b {
		if a > b {
			a = a - b
		} else {
			b = b - a
		}
	}
	return a
}

func optimisedEuclideanAlgo(x, y int) int {
	if y == 0 {
		return x
	}
	return optimisedEuclideanAlgo(y, x%y)
}

// gcd(a,b)=gcd(b,amodb)
func gcditerative(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}
