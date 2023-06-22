package main

import (
	"fmt"
	"strings"
)

func compareString(str1, str2 string) {

	if strings.Contains(str1, str2) {
		fmt.Println(str2)
	}
}

func swap(a, b *int) {
	temp := *a
	*a = *b
	*b = temp
}

func getMinMax(number ...*int) (min int, max int) {
	min = *number[0]
	max = *number[0]

	for _, num := range number {
		if *num < min {
			min = *num
		}
		if *num > max {
			max = *num
		}
	}
	return min, max
}

func main() {
	compareString("AKASHI", "AKA")

	a := 10
	b := 20
	fmt.Printf("Sebelum Swap, a = %v b = %v \n", a, b)
	swap(&a, &b)
	fmt.Printf("Sesudah Swap, a = %v b = %v\n", a, b)

	var a1, a2, a3, a4, a5, a6, min, max int
	fmt.Scan(&a1)
	fmt.Scan(&a2)
	fmt.Scan(&a3)
	fmt.Scan(&a4)
	fmt.Scan(&a5)
	fmt.Scan(&a6)
	min, max = getMinMax(&a1, &a2, &a3, &a4, &a5, &a6)
	fmt.Printf("Nilai min = %v max = %v", min, max)
}
