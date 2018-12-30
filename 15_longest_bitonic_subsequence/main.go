package main

import (
	"fmt"
)

/*
Given an array arr[0 … n-1] containing n positive integers, a subsequence of arr[] is called Bitonic if it is first
increasing, then decreasing. Write a function that takes an array as argument and returns the length of the longest
bitonic subsequence.
A sequence, sorted in increasing order is considered Bitonic with the decreasing part as empty.
Similarly, decreasing order sequence is considered Bitonic with the increasing part as empty.

Examples:
Input arr[] = {1, 11, 2, 10, 4, 5, 2, 1};
Output: 6 (A Longest Bitonic Subsequence of length 6 is 1, 2, 10, 4, 2, 1)

Input arr[] = {12, 11, 40, 5, 3, 1}
Output: 5 (A Longest Bitonic Subsequence of length 5 is 12, 11, 5, 3, 1)

Input arr[] = {80, 60, 30, 40, 20, 10}
Output: 5 (A Longest Bitonic Subsequence of length 5 is 80, 60, 30, 20, 10)

*/

var (
	input1 = []int{1, 11, 2, 10, 4, 5, 2, 1}
	input2 = []int{12, 11, 40, 5, 3, 1}
	input3 = []int{80, 60, 30, 40, 20, 10}
)

func main() {
	fmt.Printf("lbs of %v, is %d\n", input1, lbs(input1))
	fmt.Printf("lbs of %v, is %d\n", input2, lbs(input2))
	fmt.Printf("lbs of %v, is %d\n", input3, lbs(input3))

	fmt.Printf("dp lbs of %v, is %d\n", input1, lbs_dp(input1))
	fmt.Printf("dp lbs of %v, is %d\n", input2, lbs_dp(input2))
	fmt.Printf("dp lbs of %v, is %d\n", input3, lbs_dp(input3))
}

func lis(input []int, end int, maxInsc *int) int {
	if end == 0 {
		if *maxInsc < 1 {
			*maxInsc = 1
		}
		return 1
	}

	currMax := 1
	for i := 0; i < end; i++ {
		res := lis(input, i, maxInsc)
		if input[i] < input[end] && res+1 > currMax {
			currMax = res + 1
		}
	}

	if currMax > *maxInsc {
		*maxInsc = currMax
	}

	return currMax
}

func lds(input []int, start, end int, maxDesc *int) int {
	if start == end {
		if *maxDesc < 1 {
			*maxDesc = 1
		}
		return 1
	}

	currMax := 1
	for i := start; i < end; i++ {
		res := lds(input, start, i, maxDesc)
		if input[i] > input[end] && res+1 > currMax {
			currMax = res + 1
		}
	}

	if currMax > *maxDesc {
		*maxDesc = currMax
	}

	return currMax
}

func lbs(input []int) int {
	currMax := 1
	for i := 0; i < len(input); i++ {
		maxInsc := new(int)
		lis(input, i, maxInsc)
		maxDesc := new(int)
		lds(input, i, len(input)-1, maxDesc)

		if currMax < *maxInsc+*maxDesc-1 {
			currMax = *maxInsc + *maxDesc - 1
		}
	}

	return currMax
}

func lis_dp(input []int) []int {
	dpTable := make([]int, len(input))
	for i := 0; i < len(input); i++ {
		dpTable[i] = 1
	}

	for i := 1; i < len(input); i++ {
		for j := 0; j < i; j++ {
			if input[j] < input[i] && dpTable[j]+1 > dpTable[i] {
				dpTable[i] = dpTable[j] + 1
			}
		}
	}

	return dpTable
}

func lds_dp(input []int) []int {
	dpTable := make([]int, len(input))
	for i := 0; i < len(input); i++ {
		dpTable[i] = 1
	}

	for i := 1; i < len(input); i++ {
		for j := 0; j < i; j++ {
			if input[j] > input[i] && dpTable[j]+1 > dpTable[i] {
				dpTable[i] = dpTable[j] + 1
			}
		}
	}

	return dpTable
}

func lbs_dp(input []int) int {
	lisTable := lis_dp(input)
	ldsTable := lds_dp(input)
	currMax := 1
	for i := 0; i < len(input); i++ {
		if currMax < lisTable[i]+ldsTable[i]-1 {
			currMax = lisTable[i] + ldsTable[i] - 1
		}
	}

	return currMax
}
