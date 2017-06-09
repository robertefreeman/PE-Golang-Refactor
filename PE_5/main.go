/*
Project Euler Problem #5 - Go Refactor

2520 is the smallest number that can be divided by each of the numbers from 1 to 10 without any remainder.

What is the smallest positive number that is evenly divisible by all of the numbers from 1 to 20?
*/

package main

import (
	"fmt"
	"time"
)

func main() {

	var sol1, sol2 int

	// Initilize benchmark timer
	start := time.Now()

	//interate over solution x times
	for i := 0; i < 1; i++ {
		// Call finder function
		sol1 = findSmall1()
	}
	// Print first solution
	fmt.Printf("The smallest positive number evenly divisible by 1-20 is: %v\n", sol1)

	// Stop timer and print outcome of first solution
	elapsed := time.Since(start)
	fmt.Printf("Solution 5 option 1 took %s\n", elapsed)

	// Initilize benchmark timer for second option
	start = time.Now()

	//interate over solution x times
	for i := 0; i < 1; i++ {
		// Call second finder function
		sol2 = findSmall2()
	}
	// Print second solution
	fmt.Printf("The smallest positive number evenly divisible by 1-20 is: %v\n", sol2)

	// Stop timer and print outcome of second solution
	elapsed = time.Since(start)
	fmt.Printf("Solution 5 option 2 took %s", elapsed)
}

func findSmall1() int {
	poss := []int{11,12,13,14,15,16,17,18,19,20}
	smallNum := 2521
	for {
		for _, x := range poss {
			if smallNum % x != 0 {
				break
			}
			if x == 20 {
				return smallNum
			}
		}
		smallNum++
	}
}

func findSmall2() int {
	x := 2521

	for {
		if x%11==0 && x%12==0 && x%13==0 && x%14==0 && x%15==0 && x%16==0 && x%17==0 && x%18==0 && x%19==0 && x%20==0 {
			return x
		}
		x++
	}
}


