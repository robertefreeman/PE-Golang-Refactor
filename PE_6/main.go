/*
Project Euler Problem #6 - Go Refactor

The sum of the squares of the first ten natural numbers is,

12 + 22 + ... + 102 = 385
The square of the sum of the first ten natural numbers is,

(1 + 2 + ... + 10)2 = 552 = 3025
Hence the difference between the sum of the squares of the first ten natural numbers and the square of the sum is 3025 − 385 = 2640.

Find the difference between the sum of the squares of the first one hundred natural numbers and the square of the sum.

*/

package main

import (
	"fmt"
	//"strconv"
	//"math/big"
	"time"
)

func main() {

	var sol int

	// Initilize benchmark timer
	start := time.Now()

	//interate over solution x times
	for i := 0; i < 1; i++ {
		// Call finder function
		sol = ssDelta(100)
	}
	// Print solution
	fmt.Printf("The solution to Problem #6 is: %v\n", sol)

	// Stop timer and print time elapsed for solution
	elapsed := time.Since(start)
	fmt.Printf("Solution 5 option 1 took %s\n", elapsed)

}

func ssDelta(x int) int {
	y := sumSqaures(x)
	z := sqSum(x)

	return z - y

}

func sumSqaures(x int) int {
	total := 0
	for i := 1; i<=x; i++ {
		total += i*i
	}
	return total
}

func sqSum(x int) int {
	total := 0
	for i := 1; i<=x; i++ {
		total += i
	}
	return total * total
}
