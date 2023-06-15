package main

import (
	"fmt"
	"strings"
)

func lptabung() {

	t := 20.0
	r := 4.0

	luas := 2 * 3.14 * r * (r + t)
	fmt.Println(luas)
}

func nGrade() {

	n := 80

	switch {
	case n >= 80 && n <= 100:
		fmt.Println("Nilai A")
	case n >= 65:
		fmt.Println("Nilai A")
	case n >= 50:
		fmt.Println("Nilai A")
	case n >= 35:
		fmt.Println("Nilai A")
	case n >= 0:
		fmt.Println("Nilai A")
	default:
		fmt.Println("Nilai Invalid")
	}
}

func nFaktor() {
	n := 20
	result := []int{}

	for i := 1; i <= n; i++ {
		if n%i == 0 {
			result = append(result, i)
		}
	}

	fmt.Println(result)
}

func nPrima() {
	n := 100
	var logic bool

	if n < 2 {
		logic = false
	}

	for i := 2; i < n/2; i++ {
		fmt.Println(i)
		if n%i == 0 {
			logic = false
			break
		}
		logic = true
	}

	fmt.Println(logic)
}

func palindrome() {
	word := "katak"
	chars := strings.Split(word, "")
	reverse := ""
	var same bool

	for i := len(word) - 1; i >= 0; i-- {
		reverse += chars[i]
	}

	if reverse == word {
		same = true
	}
	same = false

	fmt.Println(same)
}

func exponential() {
	n := 2
	pangkat := 3
	result := 1

	for i := 1; i <= pangkat; i++ {
		result *= n
	}
	fmt.Println(result)
}

func printAsteriks() {
	n := 3

	for i := 1; i <= n; i++ {
		for j := 1; j <= i; j++ {
			fmt.Print("*")
		}
		fmt.Println("")
	}
}

func cetakTabelPerkalian() {
	n := 3

	for i := 1; i <= n; i++ {
		for j := 1; j <= n; j++ {
			fmt.Printf("%4d", i*j)
		}
		fmt.Println(" ")
	}
}

func main() {

	nGrade()
	nFaktor()
	nPrima()
	palindrome()
	exponential()
	printAsteriks()
	cetakTabelPerkalian()
}
