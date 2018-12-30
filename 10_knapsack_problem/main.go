package main

import "fmt"

/*
Given weights and values of n items, put these items in a knapsack of capacity W to get the maximum
total value in the knapsack. In other words, given two integer arrays val[0..n-1] and wt[0..n-1]
which represent values and weights associated with n items respectively.
Also given an integer W which represents knapsack capacity, find out the maximum value subset of val[]
such that sum of the weights of this subset is smaller than or equal to W. You cannot break an item,
either pick the complete item, or don’t pick it (0-1 property).

*/

var (
	vals = []int{60, 100, 120}
	wt   = []int{10, 20, 30}
)

func main() {
	// without repetition
	fmt.Printf("Knapsack soltion = %d, values = %v, weigths = %v, weigth = %d\n",
		knapsack(vals, wt, 50, 0), vals, wt, 50)
}

func knapsack(vals []int, wts []int, wt, i int) int {
	if wt <= 0 || i >= len(vals) {
		return 0
	}
	if wts[i] > wt {
		return knapsack(vals, wts, wt, i+1)
	}
	// max of (include current object, exclude current object)
	return max(vals[i]+knapsack(vals, wts, wt-wts[i], i+1), knapsack(vals, wts, wt, i+1))
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
