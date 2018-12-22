package main

import "fmt"

/*
Following are common definition of Binomial Coefficients.

A binomial coefficient C(n, k) can be defined as the coefficient of X^k in the expansion of (1 + X)^n.
A binomial coefficient C(n, k) also gives the number of ways, disregarding order, that k objects can be
chosen from among n objects; more formally, the number of k-element subsets (or k-combinations) of an n-element set.

   C(n, k) = C(n-1, k-1) + C(n-1, k)
   C(n, 0) = C(n, n) = 1

*/

type nonDP struct {
}

type dp struct {
}

func main() {
	var nonDP *nonDP

	// without repetition
	fmt.Printf("Binomial Coefficients of n = %d, k = %d, equals = %d\n",
		4, 2, nonDP.BC(5, 2))

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

func (nonDP *nonDP) BC(n, k int) int {
	if k == 0 || n == k {
		return 1
	}

	return nonDP.BC(n-1, k-1) + nonDP.BC(n-1, k)
}
