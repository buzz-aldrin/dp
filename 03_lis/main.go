package main

import "fmt"

/*
The Longest Increasing Subsequence (LIS) problem is to find the length of the longest subsequence of a given sequence
such that all elements of the subsequence are sorted in increasing order. For example, the length of LIS for
{10, 22, 9, 33, 21, 50, 41, 60, 80} is 6 and LIS is {10, 22, 33, 50, 60, 80}.
						LIS(4)
				/			|		\
			LIS(3)		LIS(2)		LIS(1)
			/	\			|
		LIS(2)	LIS(1)	LIS(1)
		/
	LIS(1)

			10	22	29	24	25	50	41	60	80
	LIS		1	2	3	3	4	5	5	6	7

Optimal Substructure:
Let arr[0..n-1] be the input array and L(i) be the length of the LIS ending at index i such that arr[i] is the last
element of the LIS. Then, L(i) can be recursively written as:
L(i) = 1 + max( L(j) ) where 0 < j < i and arr[j] < arr[i]; or
L(i) = 1, if no such j exists.
To find the LIS for a given array, we need to return max(L(i)) where 0 < i < n.
*/

var (
	input1 = []int{10, 22, 9, 33, 21, 50, 41, 60, 80, 6}
	input2 = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	input3 = []int{9, 8, 7, 6, 5, 4, 3, 2, 1}
)

// LISLoop lis loop logic
func LISRecursive(input []int, m int, max *int) int {
	// calculate LIS for current case
	// recursively call for subsequent cases
	if m == 0 {
		return 1
	}

	currLIS := 1
	for i := 0; i < m; i++ {
		res := LISRecursive(input, i, max)
		if input[i] < input[m] && res+1 > currLIS {
			currLIS = res + 1
		}
	}
	if *max < currLIS {
		*max = currLIS
	}

	return currLIS
}

func LISDP(input []int) int {
	if len(input) <= 1 {
		return 1
	}

	lisSlice := make([]int, len(input))
	for j := 1; j < len(input); j++ {
		for i := 0; i < j; i++ {
			if input[i] < input[j] {
				lisSlice[j] = max(lisSlice[i]+1, lisSlice[j])
			}
		}
	}

	return maxElem(lisSlice) + 1
}

func maxElem(lisSlice []int) int {
	max := 0
	for i := 0; i < len(lisSlice); i++ {
		if max < lisSlice[i] {
			max = lisSlice[i]
		}
	}
	return max
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func main() {
	fmt.Printf("LIS of %+v is %d\n", input1, LISDP(input1))
	fmt.Printf("LIS of %+v is %d\n", input2, LISDP(input2))
	fmt.Printf("LIS of %+v is %d\n", input3, LISDP(input3))

	lis := new(int)
	LISRecursive(input1, len(input1)-1, lis)
	fmt.Printf("LIS of %+v is %d\n", input1, *lis)
	lis = new(int)
	LISRecursive(input2, len(input2)-1, lis)
	fmt.Printf("LIS of %+v is %d\n", input2, *lis)
	lis = new(int)
	LISRecursive(input3, len(input3)-1, lis)
	fmt.Printf("LIS of %+v is %d\n", input3, *lis)
}
