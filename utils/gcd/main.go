package gcd

import "fmt"

// GCD stands for Greatest Common divisor
func GCD(a, b int) (gd int) {
	for a%b != 0 {
		r := a % b
		if a > b {
			a = b
			b = r
		} else {
			b = a
			a = r
		}
		fmt.Println(a, b, a%b)
	}

	return b
}
