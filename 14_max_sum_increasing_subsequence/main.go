package main

import "fmt"

/*
Given an array of n positive integers. Write a program to find the sum of maximum sum subsequence of the given array
such that the integers in the subsequence are sorted in increasing order.
Examples:
	if input is {1, 101, 2, 3, 100, 4, 5}, then output should be 106 (1 + 2 + 3 + 100),
	if the input array is {3, 4, 5, 10}, then output should be 22 (3 + 4 + 5 + 10)
	if the input array is {10, 5, 4, 3}, then output should be 10
*/

var (
	input1 = []int{1, 101, 2, 3, 100, 4, 5}
	input2 = []int{3, 4, 5, 10}
	input3 = []int{10, 5, 4, 3}
)

func main() {
	max := new(int)
	msis(input1, len(input1)-1, max)
	fmt.Printf("msis of %v, is %d\n", input1, *max)
	*max = 0
	msis(input2, len(input2)-1, max)
	fmt.Printf("msis of %v, is %d\n", input2, *max)
	*max = 0
	msis(input3, len(input3)-1, max)
	fmt.Printf("msis of %v, is %d\n", input3, *max)

	fmt.Printf("dp msis of %v, is %d\n", input1, msis_dp(input1))
	fmt.Printf("dp msis of %v, is %d\n", input2, msis_dp(input2))
	fmt.Printf("dp msis of %v, is %d\n", input3, msis_dp(input3))
}

func msis(input []int, end int, max *int) int {
	if end == 0 {
		if *max < input[0] {
			*max = input[0]
		}
		return input[0]
	}

	var currMax int
	for i := 0; i < end; i++ {
		temp := msis(input, i, max)
		if input[i] < input[end] && temp+input[end] > currMax {
			currMax = input[end] + temp
		}
	}

	if *max < currMax {
		*max = currMax
	}
	return currMax
}

func msis_dp(input []int) int {
	dpTable := make([]int, len(input))
	for i := 0; i < len(input); i++ {
		dpTable[i] = input[i]
	}

	for i := 1; i < len(input); i++ {
		for j := 0; j < i; j++ {
			if input[j] < input[i] && dpTable[j]+input[i] > dpTable[i] {
				dpTable[i] = dpTable[j] + input[i]
			}
		}
	}

	return maxElem(dpTable)
}

func maxElem(input []int) int {
	max := 0
	for i := 0; i < len(input); i++ {
		if max < input[i] {
			max = input[i]
		}
	}
	return max
}
