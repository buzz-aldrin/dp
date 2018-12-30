package main

import "fmt"

/*
Longest Common Subsequence - https://www.geeksforgeeks.org/longest-common-subsequence-dp-4/
It is a classic computer science problem, the basis of diff (a file comparison program that outputs the differences
between two files), and has applications in bio-informatics.

Examples:
LCS for input Sequences “ABCDGH” and “AEDFHR” is “ADH” of length 3.
LCS for input Sequences “AGGTAB” and “GXTXAYB” is “GTAB” of length 4.

                            lcs("AXYT", "AYZX")
                           /                 \
             lcs("AXY", "AYZX")            lcs("AXYT", "AYZ")
             /                \              /               \
    lcs("AX", "AYZX") lcs("AXY", "AYZ")   lcs("AXY", "AYZ") lcs("AXYT", "AY")

		A	G	G	T	A	B
	0	0	0	0	0	0	0
G	0	0	1	1	1	1	1
X	0	0	1	1	1	1	1
T	0	0	1	1	2	2	2
X	0	0	1	1	2	2	2
A	0	1	2	2	2	3	3
Y	0	1	2	2	2	3	3
B	0	1	2	2	2	3	4

*/

var (
	str1 = []rune{'G', 'X', 'T', 'X', 'A', 'Y', 'B'}
	str2 = []rune{'A', 'G', 'G', 'T', 'A', 'B'}
)

func main() {
	fmt.Printf("LCS of %c and %c = %d\n", str1, str2, LCSDP(str1, str2))

	fmt.Printf("LCS of %c and %c = %d\n", str1, str2, LCSRecursive(str1, str2, 0, 0))
}

func LCSRecursive(str1, str2 []rune, x, y int) int {
	if x >= len(str1) || y >= len(str2) {
		return 0
	}

	if str1[x] == str2[y] {
		return 1 + LCSRecursive(str1, str2, x+1, y+1)
	}
	return max(LCSRecursive(str1, str2, x+1, y), LCSRecursive(str1, str2, x, y+1))

}

// LCS logic
func LCSDP(str1, str2 []rune) int {
	dpTable := initDPTable(len(str1)+1, len(str2)+1)

	for i := 1; i <= len(str1); i++ {
		for j := 1; j <= len(str2); j++ {
			if str1[i-1] == str2[j-1] {
				dpTable[i][j] = 1 + dpTable[i-1][j-1]
			} else {
				dpTable[i][j] = max(dpTable[i][j-1], dpTable[i-1][j])
			}
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

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
