package main

import "fmt"

func main() {
	maxNumber := 20

	for num := 2; num <= maxNumber; num++ {
		isPrime := true

		for divisor := 2; divisor < num; divisor++ {
			if num%divisor == 0 {
				isPrime = false
				break
			}
		}
		if isPrime {
			fmt.Println(num)
		}
	}
}
