package main

import "fmt"

func main() {
	var n int = 10
	var k int = 0
	var divisor [5]int = [5]int{}

	for i := 1; i <= n/2; i++ {
		if n%i == 0 {
			divisor[k] = i
			k += 1
		}
	}

	divisor[k] = n
	fmt.Printf("Divisors of %d are: %d", n, divisor)
}

// I need to make a dynamic array for this to make it more efficient. Will update this soon.