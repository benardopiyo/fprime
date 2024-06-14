package main

import (
	"os"
	"github.com/01-edu/z01"
)

// atoi converts a string to an integer without using strconv.
func atoi(s string) (int, bool) {
	if s == "" {
		return 0, false
	}
	result := 0
	for _, r := range s {
		if r < '0' || r > '9' {
			return 0, false
		}
		result = result*10 + int(r-'0')
	}
	return result, true
}

// primeFactors returns a slice of prime factors of the given number.
func primeFactors(n int) []int {
	if n < 2 {
		return []int{}
	}
	factors := []int{}
	for i := 2; i*i <= n; i++ {
		for n%i == 0 {
			factors = append(factors, i)
			n /= i
		}
	}
	if n > 1 {
		factors = append(factors, n)
	}
	return factors
}

// printFactors prints the prime factors in the required format.
func printFactors(factors []int) {
	for i, factor := range factors {
		if i > 0 {
			z01.PrintRune('*')
		}
		for _, r := range itoa(factor) {
			z01.PrintRune(r)
		}
	}
	z01.PrintRune('\n')
}

// itoa converts an integer to a string without using strconv.
func itoa(n int) string {
	if n == 0 {
		return "0"
	}
	result := ""
	for n > 0 {
		result = string(rune('0'+(n%10))) + result
		n /= 10
	}
	return result
}

func main() {
	if len(os.Args) != 2 {
		return
	}

	num, ok := atoi(os.Args[1])
	if !ok || num < 2 {
		return
	}

	factors := primeFactors(num)
	if len(factors) > 0 {
		printFactors(factors)
	}
}
