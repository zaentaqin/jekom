package main

import (
	"fmt"
	"strings"
)

func arrayMerge() {
	str1 := []string{"muhammad", "zaenal"}
	str2 := []string{"muhammad", "muttaqin"}
	merge := []string{}

	//merge all element str1
	merge = append(merge, str1...)

	//iterate str 2 and add element not the same str 1
	var found bool
	for _, str := range str2 {
		found = false
		for _, val := range str1 {
			if str == val {
				found = true
				break
			}
		}
		if !found {
			merge = append(merge, str)
		}

	}
	fmt.Println("Array merge :", merge)

}

func munculSekali() {
	angka := "2211334"
	separate := strings.Split(angka, "")

	var isDuplicate bool
	var result []string

	for i := 0; i < len(separate); i++ {
		isDuplicate = false
		for j := 0; j < len(separate); j++ {
			if i != j && separate[i] == separate[j] {
				isDuplicate = true
				break
			}
		}

		if !isDuplicate {
			result = append(result, separate[i])
		}
	}

	fmt.Println("Number that appears once :", result)
}

func pairSum() {
	arr := []int{1, 2, 3, 4, 6}
	target := 6

	var result []int

	for i := 0; i < len(arr)-1; i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i]+arr[j] == target {
				result = []int{i, j}
				break
			}
		}
	}

	if len(result) == 2 {
		fmt.Println("Pair found index :", result[0], result[1])
	} else {
		fmt.Println("No pair found.")
	}
}

func main() {
	arrayMerge()
	munculSekali()
	pairSum()
}
