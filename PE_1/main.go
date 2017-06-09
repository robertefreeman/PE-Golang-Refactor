/*
Project Euler Problem #1 - Go Refactor

If we list all the natural numbers below 10 that are multiples of 3 or 5, we get 3, 5, 6 and 9. The sum of these multiples is 23.

Find the sum of all the multiples of 3 or 5 below 1000.

 */

package main

import (
	"fmt"
	//"math/big"
	"time"
)

func main() {
	// Initilize benchmark timer
	start := time.Now()

	// Create solution and call function
	counter := multiples(1000)
	// Print solution
	fmt.Println(counter)

	// Stop timer and print outcome
	elapsed := time.Since(start)
	fmt.Printf("Solution 20 took %s", elapsed)
}

func multiples(x int) int {
	counter := 0
	for i := 0; i < 1000; i++ {
		if i%3 == 0 {
			counter += i
		} else if i%5 == 0 {
			counter += i
		}
	}
	return counter
}