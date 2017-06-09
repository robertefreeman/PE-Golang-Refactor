/*
Project Euler Problem #4 - Go Refactor

A palindromic number reads the same both ways. The largest palindrome made from the product of two 2-digit numbers is 9009 = 91 × 99.

Find the largest palindrome made from the product of two 3-digit numbers.
*/

package main

import (
	"fmt"
	"strconv"
	//"math/big"
	"time"
)

func main() {
	// start benchmark timer
	start := time.Now()

	var sol int

	//interate over solution x times for benchmarking
	for i := 0; i < 1; i++ {
		//Call large Palindrome finder function
		sol = findLarge()
	}
	// Print solution
	fmt.Printf("The largest palindrome made from the product of two 3-digit numbers is: %v\n", sol)

	// Stop timer and print outcome
	elapsed := time.Since(start)
	fmt.Printf("Solution 25 took %s", elapsed)
}

func isPal(x int) bool {
	str := strconv.Itoa(x)
	var reverse string
	for i := len(str) - 1; i >= 0; i-- {
		reverse += string(str[i])
	}
	if str == reverse {
		return true
	}
	return false
}

func findLarge() int {
	larPal := 1
	x := 999
	y := 999
	for x > 99 {
		for y >= x {
			if isPal(x * y) {
				if (x * y) > larPal {
					larPal = x * y
				}
			}
			y --
		}
		x --
	}
	return larPal
}
