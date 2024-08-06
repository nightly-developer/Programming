//go:build ignore

package main

func isPalindrome(x int) int {
	result := 1
	for i := 0; x > 0; i++ {
		var div int = int(x / 10)
		var rem int = int(x % 10)
		result *= rem * 10
		x = div
	}
	return result
}
