/*
Project Euler Problem #4 - Go Refactor

2520 is the smallest number that can be divided by each of the numbers from 1 to 10 without any remainder.

What is the smallest positive number that is evenly divisible by all of the numbers from 1 to 20?
*/

package main

import (
	"fmt"
	"strconv"
	//"math/big"
	"time"
)

func main() {
	// Initilize benchmark timer
	start := time.Now()

	var sol int

	//interate over solution x times
	for i := 0; i < 10000; i++ {
		//Call large Palindrome finder function
		sol = findLarge()
	}
	// Print solution
	fmt.Printf("The largest palindrome made from the product of two 3-digit numbers is: %v\n", sol)

	// Stop timer and print outcome
	elapsed := time.Since(start)
	fmt.Printf("Solution 25 took %s", elapsed)
}

/*
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
*/

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
	x:= 999
	y:=999
	for x > 99 {
		for y >= x {
			if isPal(x*y) {
				if (x * y)>larPal {
					larPal = x*y
				}
			}
			y -= 1
		}
		x -= 1
	}
	return larPal
}
