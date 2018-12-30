package main

import (
	"fmt"
	"math"
)

/*
Given a value N, if we want to make change for N cents, and we have infinite supply of each of S = { S1, S2, .. , Sm}
valued coins, how many ways can we make the change? The order of coins doesn’t matter.

For example, for N = 4 and S = {1,2,3}, there are four solutions: {1,1,1,1},{1,1,2},{2,2},{1,3}.
So output should be 4.
For N = 10 and S = {2, 5, 3, 6}, there are five solutions: {2,2,2,2,2}, {2,2,3,3}, {2,2,6}, {2,3,5} and {5,5}.
So the output should be 5.


*/

func main() {
	// without repetition
	fmt.Printf("min = %d, floors = %d\n",
		eggDrop(10, 2), 100)

}

func eggDrop(floors int, eggs int) int {
	if eggs == 1 {
		return floors
	}

	if floors <= 1 {
		return floors
	}

	min := math.MinInt32
	for i := 1; i <= floors; i++ {
		res := max(eggDrop(i-1, eggs-1), eggDrop(floors-i, eggs))
		if res < min {
			min = res
		}
	}
	return min + 1
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
