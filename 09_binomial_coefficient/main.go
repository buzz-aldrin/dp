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

func main() {
	fmt.Printf("Binomial Coefficients of n = %d, k = %d, equals = %d\n", 4, 2, bc(5, 2))
}

func bc(n, k int) int {
	if k == 0 || n == k {
		return 1
	}

	return bc(n-1, k-1) + bc(n-1, k)
}
