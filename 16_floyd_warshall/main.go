package main

import (
	"fmt"
	"math"
)

/*
Pick a node intermediate node k such that 0 <= k < no of nodes
Now for each source and destination pair go through k and check if distance is less in case we are moving through k
instead of jumping directly from origin to destination. If true update

The Floyd Warshall Algorithm is for solving the All Pairs Shortest Path problem. The problem is to find shortest
distances between every pair of vertices in a given edge weighted directed Graph.

Example:
	Input:
		   graph[][] = { {0,   5,  INF, 10},
						{INF,  0,  3,  INF},
						{INF, INF, 0,   1},
						{INF, INF, INF, 0} }
	which represents the following graph
				 10
		   (0)------->(3)
			|         /|\
		  5 |          |
			|          | 1
		   \|/         |
		   (1)------->(2)
				3
	Note that the value of graph[i][j] is 0 if i is equal to j
	And graph[i][j] is INF (infinite) if there is no edge from vertex i to j.

	Output:
	Shortest distance matrix
		  0      5      8      9
		INF      0      3      4
		INF    INF      0      1
		INF    INF    INF      0

Floyd Warshall Algorithm
We initialize the solution matrix same as the input graph matrix as a first step. Then we update the solution matrix by
considering all vertices as an intermediate vertex. The idea is to one by one pick all vertices and updates all shortest
paths which include the picked vertex as an intermediate vertex in the shortest path. When we pick vertex number k as an
intermediate vertex, we already have considered vertices {0, 1, 2, .. k-1} as intermediate vertices. For every pair
(i, j) of the source and destination vertices respectively, there are two possible cases.
1) k is not an intermediate vertex in shortest path from i to j. We keep the value of dist[i][j] as it is.
2) k is an intermediate vertex in shortest path from i to j. We update the value of dist[i][j] as
   dist[i][k] + dist[k][j] if dist[i][j] > dist[i][k] + dist[k][j]
*/

var (
	inf    = math.MaxInt32
	input1 = [][]int{
		{0, 5, inf, 10},
		{inf, 0, 3, inf},
		{inf, inf, 0, 1},
		{inf, inf, inf, 0},
	}
	input2 = [][]int{
		{0, 3, 6, 15},
		{inf, 0, -2, inf},
		{inf, inf, 0, 2},
		{1, inf, inf, 0},
	}
)

func main() {
	//printFW(input1)
	printFW(input2)
}

func printFW(input [][]int) {
	fmt.Println("for: ")
	printTable(input)
	fmt.Println("all pair shortest path: ")
	printTable(fw(input))
}

func fw(input [][]int) (dist [][]int) {
	dist = distTable(input)
	// if k is an intermediate vertex in shortest path from i to j.
	// We update the value of dist[i][j] as dist[i][k] + dist[k][j]
	// if dist[i][j] > dist[i][k] + dist[k][j]

	// Pick all intermediate vertices between source and destination
	for k := 0; k < len(input); k++ {
		// Pick all vertices as source one by one
		for i := 0; i < len(input); i++ {
			// Pick all vertices as destination for the above picked source
			for j := 0; j < len(input[i]); j++ {
				if dist[i][j] > dist[i][k]+dist[k][j] {
					dist[i][j] = dist[i][k] + dist[k][j]
				}
			}
		}
	}

	return
}

func distTable(input [][]int) (dist [][]int) {
	dist = make([][]int, len(input))
	for i := 0; i < len(input); i++ {
		dist[i] = make([]int, len(input[i]))
		for j := 0; j < len(input[i]); j++ {
			dist[i][j] = input[i][j]
		}
	}
	return
}

func printTable(input [][]int) {
	for i := 0; i < len(input); i++ {
		for j := 0; j < len(input[i]); j++ {
			if input[i][j] == inf {
				fmt.Printf("n ")
			} else {
				fmt.Printf("%d ", input[i][j])
			}
		}
		fmt.Println()
	}
	return
}
