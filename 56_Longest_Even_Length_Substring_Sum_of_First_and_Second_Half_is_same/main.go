// Longest Even Length Substring such that Sum of First and Second Half is same
package main

var (
	str = []int{1, 5, 3, 8, 0, 2, 3}
)

func main() {

}

// all possible sub strings of even length in first half and second half

// Select the largest possible even sub string
// if len(ip)%2 == 0 then len(ip)/2 else len(s)/2 - 1
func ELS(ip []int) {
	lps := len(ip) / 2
	for i := lps; i > 0; i = i - 2 {
		for l := 0; l < lps; l++ {

		}
	}

}
