package main

import "fmt"

/*
Edit Distance - https://www.geeksforgeeks.org/edit-distance-dp-5/
Given two strings str1 and str2 and below operations that can performed on str1. Find minimum number of edits
(operations) required to convert ‘str1’ into ‘str2’.
	a. Insert
	b. Remove
	c. Replace
All of the above operations are of equal cost.
Examples:
	Input:   str1 = "geek", str2 = "geeks"
	Output:  1
	We can convert str1 into str2 by inserting a 's'.

	Input:   str1 = "cat", str2 = "cut"
	Output:  1
	We can convert str1 into str2 by replacing 'a' with 'u'.

			s	u	n	d	a	y
		0	0	0	0	0	0	0
	s	0	0	1	1	1	1	1
	a	0	1	1	2	2	1	2
	t	0	1	2	2	3	2	2
	u	0	1	1	2	3	3	3
	r	0	1	2	2	3	4	4
	d	0	1	2	3	2	3	4
	a	0	1	2	3	3	2	3
	y	0	1	2	3	4	3	2
*/

var (
	str1 = []rune{'s', 'a', 't', 'u', 'r', 'd', 'a', 'y'}
	str2 = []rune{'s', 'u', 'n', 'd', 'a', 'y'}
	//str2 = []rune{'1', '1', '1', '1', '1', '1'}
)

func main() {

	fmt.Printf("EDRecursive for %c, %c = %d\n", str1, str2, EDRecursive(str1, str2, len(str1)-1, len(str2)-1))

	fmt.Printf("EDDynamic for %c, %c = %d\n", str1, str2, EDDP(str1, str2))

}

func EDRecursive(str1, str2 []rune, x, y int) int {
	if x == 0 {
		return y
	}

	if y == 0 {
		return x
	}

	if str1[x] == str2[y] {
		return EDRecursive(str1, str2, x-1, y-1)
	}
	return 1 + min(
		EDRecursive(str1, str2, x, y-1),   // Insert
		EDRecursive(str1, str2, x-1, y),   // Remove
		EDRecursive(str1, str2, x-1, y-1), // Replace
	)
}

func EDDP(str1, str2 []rune) int {
	dpTable := initDPTable(len(str1)+1, len(str2)+1)

	for i := 0; i <= len(str1); i++ {
		for j := 0; j <= len(str2); j++ {
			if i == 0 {
				dpTable[i][j] = j
				continue
			}
			if j == 0 {
				dpTable[i][j] = i
				continue
			}
			if str1[i-1] == str2[j-1] {
				dpTable[i][j] = dpTable[i-1][j-1]
				continue
			}
			dpTable[i][j] = 1 + min(dpTable[i][j-1], dpTable[i-1][j], dpTable[i-1][j-1])
		}
	}

	return dpTable[len(str1)][len(str2)]
}

func initDPTable(l1, l2 int) [][]int {
	dpTable := make([][]int, l1)
	for i := 0; i < len(dpTable); i++ {
		dpTable[i] = make([]int, l2)
	}

	return dpTable
}

func min(a, b, c int) int {
	var min int
	if a < b {
		min = a
	} else {
		min = b
	}

	if c < min {
		return c
	}
	return min
}

func printTable(dpTable [][]int) {
	fmt.Printf("    %c\n", str2)
	for i := 0; i < len(dpTable); i++ {
		if i == 0 {
			fmt.Printf("  %+v\n", dpTable[i])
		} else {
			fmt.Printf("%c %+v\n", str1[i-1], dpTable[i])
		}
	}
}
