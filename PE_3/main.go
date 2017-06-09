/*
Project Euler Problem #3 - Go Refactor

The prime factors of 13195 are 5, 7, 13 and 29.

What is the largest prime factor of the number 600851475143 ?
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

	var sol int
	//interate over solution x times
	for i := 0; i < 100000; i++ {
		//Call even Fibs function
		sol = primeFactor(600851475143)
	}

	// Print solution
	fmt.Printf("The largest prime factor of the number is: %v\n", sol)

	// Stop timer and print outcome
	elapsed := time.Since(start)
	fmt.Printf("Solution 25 took %s", elapsed)
}

func primeFactor(val int) int {
	solution := 0
	f := 3
	// Loop for odd numbers as factors
	for val != 1 {
		for val % f ==0 {
			solution = f
			val = val/f
		}
		f += 2
	}
	return solution
}
