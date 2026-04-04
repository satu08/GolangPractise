package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(evenOdd(5))
	multiplicationTable(5)
	fmt.Println(sumOfNaturalNumbers(10))
	fmt.Println(sumofSquaresOfNaturals(8))
	fmt.Println(swapTwoNumbers(4, 5))
	fmt.Println(oppositeNumberOnCubicDice(5))
	fmt.Println(nthTermInArithmeticProgression(1, 3, 10))
	fmt.Println(sumOfDigits(75))
	fmt.Println(reverseDigits(1212))
	fmt.Println(checkIfPower(10, 1001))
	fmt.Println(distanceOfTwoPoints(3, 4, 7, 7))
	fmt.Println(validTriangle(7, 10, 5))
	fmt.Println(pairCubeCount(28))
	fmt.Println(perfectNumber(6))
	fmt.Println(decimalToBinary(10))
}

func evenOdd(n int) bool {
	if n/2 == 0 {
		return true
	}
	return false
}

func multiplicationTable(n int) {
	for i := 1; i <= 10; i++ {
		fmt.Println(n, "*", i, "=", n*i)
	}
}

func sumOfNaturalNumbers(n int) int {
	return (n * (n + 1)) / 2
}

func sumofSquaresOfNaturals(n int) int {
	sum := 0
	for i := 1; i <= n; i++ {
		sum += i * i
	}
	return sum
}

func swapTwoNumbers(a, b int) (int, int) {
	// Swap a and b using temp variable
	//int temp = a;
	//a = b;
	//b = temp;
	a = a + b
	b = a - b
	a = a - b
	return a, b
}

func oppositeNumberOnCubicDice(n int) int {
	return 7 - n
}

func nthTermInArithmeticProgression(a1, a2, n int) int {
	return a1 + ((a2 - a1) * (n - 1))
}

func sumOfDigits(n int) int {
	sum := 0
	for n > 0 {
		modulo := n % 10
		sum += modulo
		n /= 10
	}
	return sum
}

func reverseDigits(n int) int {
	reverse := 0
	for n > 0 {
		modulo := n % 10
		reverse = reverse*10 + modulo
		n /= 10
	}
	return reverse
}

func checkIfPower(n, m int) bool {
	if m%n == 0 {
		return true
	}
	return false
}

func distanceOfTwoPoints(x1, y1, x2, y2 int) float64 {
	dist1 := (x2 - x1) * (x2 - x1)
	dist2 := (y2 - y1) * (y2 - y1)
	dist3 := dist1 + dist2
	return math.Sqrt(float64(dist3))

}

func validTriangle(a, b, c int) bool {
	if a+b <= c || b+c <= a || a+c <= b {
		return false
	}
	return true
}

func pairCubeCount(n int) int {
	count := 0
	for i := 1; i < n; i++ {
		for j := 1; j < n; j++ {
			sum := (i * i * i) + (j * j * j)
			if sum == n {
				count++
			}
		}
	}
	return count
}

func perfectNumber(n int) bool {
	sum := 0
	for i := 1; i < n; i++ {
		if n%i == 0 {
			sum += i
		}
	}
	return sum == n
}

func decimalToBinary(n int) string {
	bits := ""
	for n > 0 {
		mod := n % 2
		bits = fmt.Sprint(mod) + bits
		n = n / 2
	}
	return bits
}
