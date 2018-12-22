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

type nonDP struct {
}

type dp struct {
}

var (
	vals = []int{60, 100, 120}
	wt   = []int{10, 20, 30}
)

func main() {
	var nonDP *nonDP

	// without repetition
	fmt.Printf("Knapsack soltion = %d, values = %v, weigths = %v, weigth = %d\n",
		nonDP.Knapsack(vals, wt, 50, 0), vals, wt, 50)
}

func (nonDP *nonDP) Knapsack(vals []int, wts []int, wt, i int) int {
	if wt <= 0 || i >= len(vals) {
		return 0
	}
	if wts[i] > wt {
		return nonDP.Knapsack(vals, wts, wt, i+1)
	}
	// include current object
	v1 := vals[i] + nonDP.Knapsack(vals, wts, wt-wts[i], i+1)
	// exclude current object
	v2 := nonDP.Knapsack(vals, wts, wt, i+1)
	return max(v1, v2)
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

//func (dp *dp) Knapsack(vals []int, wts []int, wt int){
//	table := make([]int, len(vals))
//
//	for i := 0; i < len(table); i++{
//		tWt := wt
//		for j := i; j < len(vals); j++{
//			if tWt - wts[j] < 0{
//				continue
//			}
//			table[i] = vals[i] +
//		}
//	}
//}
