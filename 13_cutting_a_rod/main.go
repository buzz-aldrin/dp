package main

import "fmt"

/*
Given a rod of length n inches and an array of prices that contains prices of all pieces of size smaller than n.
Determine the maximum value obtainable by cutting up the rod and selling the pieces. For example, if length of
the rod is 8 and the values of different pieces are given as following, then the maximum obtainable value is 22

length   | 1   2   3   4   5   6   7   8
--------------------------------------------
price    | 1   5   8   9  10  17  17  20
*/

var (
	vals = []int{1, 5, 8, 9, 10, 17, 17, 20}
)

func main() {
	fmt.Printf("cutting rod = %v, is %d\n", vals, cuttingRod(vals, len(vals)-1))
}

func cuttingRod(vals []int, size int) int {
	var maxVal int
	for i := 0; i <= size; i++ {
		temp := max(vals[i]+cuttingRod(vals, size-i-1), cuttingRod(vals, size-i-1))
		if temp > maxVal {
			maxVal = temp
		}
	}
	return maxVal
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
