package main

import "fmt"

/*
Given a value N, if we want to make change for N cents, and we have infinite supply of each of S = { S1, S2, .. , Sm}
valued coins, how many ways can we make the change? The order of coins doesn’t matter.

For example, for N = 4 and S = {1,2,3}, there are four solutions: {1,1,1,1},{1,1,2},{2,2},{1,3}.
So output should be 4.
For N = 10 and S = {2, 5, 3, 6}, there are five solutions: {2,2,2,2,2}, {2,2,3,3}, {2,2,6}, {2,3,5} and {5,5}.
So the output should be 5.


*/

func coinChange(coins []int, change, x int) int {
	if x >= len(coins) {
		return 0
	}

	if change < 0 {
		return 0
	}

	if change-coins[x] == 0 {
		return 1
	}

	return sum(
		coinChange(coins, change, x+1),
		coinChange(coins, change-coins[x], x+1),
		coinChange(coins, change-coins[x], x),
	)
}

func sum(arr ...int) int {
	sum := 0
	for i := 0; i < len(arr); i++ {
		sum += arr[i]
	}
	return sum
}

var (
	//coins = []int{1, 2, 3}
	coins = []int{2, 5, 3, 6}
)

func main() {
	fmt.Printf("possible coin change solutions = %d, for sum = %d, coins = %d\n",
		coinChange(coins, 7, 0), 7, coins)
	//var nonDP *nonDP

	// without repetition
	//fmt.Printf("possible coin change solutions = %d, for sum = %d, coins = %d\n",
	//	nonDP.withoutRepetition(arr, 7, 0), 7, arr)
	//
	//fmt.Printf("possible coin change solutions = %d, for sum = %d, coins = %d\n",
	//	nonDP.coinChange(coins, 4, 0), 4, coins)
	//fmt.Printf("possible coin change solutions = %d, for sum = %d, coins = %d\n",
	//	nonDP.coinChange(coins, 10, 0), 10, coins)
	//
	//var dp *dp
	//fmt.Printf("dp::possible coin change solutions = %d, for sum = %d, coins = %d\n",
	//	dp.coinChange(coins, 4), 4, coins)
	//fmt.Printf("dp::possible coin change solutions = %d, for sum = %d, coins = %d\n",
	//	dp.coinChange(coins, 10), 10, coins)

}

//func (nonDP *nonDP) coinChange(coins []int, sum int, index int) int {
//	if sum == 0 {
//		return 1
//	}
//	if sum < 0 {
//		return 0
//	}
//
//	if index >= len(coins) {
//		return 0
//	}
//	// including index in sum calculation and jumping to next index
//	s1 := nonDP.coinChange(coins, sum-coins[index], index+1)
//	// including index in sum calculation and considering same index again(coin repetition)
//	s2 := nonDP.coinChange(coins, sum-coins[index], index)
//	// excluding index from sum calculation and jump to next one
//	s3 := nonDP.coinChange(coins, sum, index+1)
//	return s1 + s2 + s3
//}
//
//// store number of solutions for each element
//// create slice of equal size as input array
//
//// let n1 and n2 be 2 numbers
//// if n1 - n2 == 0 => it is possible to make n1 using n2 => 1 solution possible
//
//// create a array equal to the sum of coins to be calculated
//// at zeroth index put 1(check above for reason)
//
////  consider each index of new table as a number and check if it is possible to make that number
//// using the number in the coin array
//func (dp *dp) coinChange(coins []int, sum int) int {
//	table := make([]int, sum+1)
//	table[0] = 1
//
//	for i := 0; i < len(coins); i++ {
//		for j := coins[i]; j <= sum; j++ {
//			table[j] += table[j-coins[i]]
//		}
//	}
//
//	return table[sum-1]
//}
//
//// array
//var arr = []int{0, -2, 2, 7}
//
//func (nonDP *nonDP) withoutRepetition(coins []int, sum int, index int) int {
//	if sum == 0 {
//		return 1
//	}
//
//	if index >= len(coins) {
//		return 0
//	}
//
//	// including index in sum calculation and jumping to next index
//	if sum > 0 {
//		return nonDP.withoutRepetition(coins, sum-coins[index], index+1) + nonDP.withoutRepetition(coins, sum, index+1)
//	}
//	return nonDP.withoutRepetition(coins, sum+coins[index], index+1) + nonDP.withoutRepetition(coins, sum, index+1)
//}
