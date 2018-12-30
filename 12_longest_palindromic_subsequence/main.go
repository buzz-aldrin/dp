package main

import "fmt"

/*
if the given sequence is “BBABCBCAB”, then the output should be 7 as “BABCBAB” is the longest palindromic
subseuqnce in it. “BBBBB” and “BBCBB” are also palindromic subsequences of the given sequence, but not the longest ones.

*/
var (
	//str = []rune{'B', 'B', 'A', 'B', 'C', 'B', 'C', 'A', 'B'}
	str = []rune{'G', 'E', 'E', 'K', 'S', 'F', 'O', 'R', 'G', 'E', 'E', 'K', 'S'}
)

func main() {
	// without repetition
	fmt.Printf("longest palindromic subsequence for = %c, is %d\n",
		str, lps(str, 0, len(str)-1))

}

func lps(str []rune, i, j int) int {
	if i == j {
		return 1
	}
	if str[i] == str[j] && i+1 == j {
		return 2
	}

	if str[i] == str[j] {
		return lps(str, i+1, j-1) + 2
	}

	return max(lps(str, i, j-1), lps(str, i+1, j))
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
