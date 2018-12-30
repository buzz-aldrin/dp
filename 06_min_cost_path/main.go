package main

import (
	"fmt"
	"math"
)

/*
Given a cost matrix cost[][] and a position (m, n) in cost[][], write a function that returns cost of minimum cost path
to reach (m, n) from (0, 0). Each cell of the matrix represents a cost to traverse through that cell. Total cost of a
path to reach (m, n) is sum of all the costs on that path (including both source and destination). You can only
traverse down, right and diagonally lower cells from a given cell,
i.e., from a given cell (i, j), cells (i+1, j), (i, j+1) and (i+1, j+1) can be traversed.
You may assume that all costs are positive integers.

1) Optimal Substructure
The path to reach (m, n) must be through one of the 3 cells: (m-1, n-1) or (m-1, n) or (m, n-1). So minimum cost to
reach (m, n) can be written as “minimum of the 3 cells plus cost[m][n]”.

minCost(m, n) = min (minCost(m-1, n-1), minCost(m-1, n), minCost(m, n-1)) + cost[m][n]

2) Overlapping Subproblems
Following is simple recursive implementation of the MCP (Minimum Cost Path) problem. The implementation simply follows
the recursive structure mentioned above.
*/

var (
	cost = [][]int{
		{1, 2, 3},
		{4, 8, 2},
		{1, 5, 3},
	}
)

func main() {
	fmt.Printf("minCostPathRecursive for %v = %d\n", cost, minCostPathRecursive(cost, 0, 0))
}

func minCostPathRecursive(input [][]int, x, y int) int {
	if x >= len(input) || y >= len(input[x]) {
		return math.MaxInt32
	}

	if x == len(input)-1 && y == len(input[x])-1 {
		return input[x][y]
	}

	return input[x][y] + min(
		minCostPathRecursive(input, x+1, y),
		minCostPathRecursive(input, x, y+1),
		minCostPathRecursive(input, x+1, y+1),
	)
}

func minCostPathDP(input [][]int, len1, len2 int) int {
	dpTable := initDPTable(len1, len2)

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
