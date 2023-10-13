package main

import "fmt"

// a series of numbers in which each number ( Fibonacci number )
//is the sum of the two preceding numbers.
//The simplest is the series 1, 1, 2, 3, 5, 8, etc.
func fibonacci(n int) int {
	if n < 2 {
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

func main() {
	fmt.Println(fibonacci(5))

}
