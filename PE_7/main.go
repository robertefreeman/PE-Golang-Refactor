/*
Project Euler Problem #7 - Go Refactor

By listing the first six prime numbers: 2, 3, 5, 7, 11, and 13, we can see that the 6th prime is 13.

What is the 10,001st prime number?

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
		sol = findPrime(10001)
	}
	// Print solution
	fmt.Printf("The solution to Problem #7 is: %v\n", sol)

	// Stop timer and print time elapsed for solution
	elapsed := time.Since(start)
	fmt.Printf("Solution 5 option 1 took %s\n", elapsed)

}
func findPrime(x int) int {
	primeCount := 0
	finalPrime := 0
	i := 1
	for {
		if isPrime(i) {
			primeCount++
			if primeCount == x {
				finalPrime = i
				return finalPrime
			}
		}
		i++

	}
	return 0
}

func isPrime(x int) bool {

	// We know 1 is not a prime number
	if x == 1 {
		return false
	}

	i := 2
	// This will loop from 2 to the sqrt(x)
	for i*i <= x {
		if x % i == 0 {
			return false
		}
		i +=1
	}
	return true
}

