package main

import (
	"fmt"
	"math"
)

func primeNumber(n int) (prime bool) {
	if n <= 1 {
		return false
	}

	for i := 2; i <= int(math.Sqrt(float64(n))); i++ {
		if n%i == 0 {
			return false
		}
	}

	return true
}

func pow(base, exponential int) int {
	result := 1

	for exponential > 0 {
		if exponential%2 == 1 {
			result *= base
		}
		base *= base
		exponential /= 2
	}
	return result
}

func main() {
	n := 123321426
	if primeNumber(n) {
		fmt.Println(n, "adalah bilangan prima")
	} else {
		fmt.Println(n, "bukan bilangan prima")
	}

	base := 2
	exponential := 30

	ress := pow(base, exponential)
	fmt.Printf("base %v, exponential %v, hasil %v \n", base, exponential, ress)
}
