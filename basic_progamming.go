package main

import (
	"fmt"
	"strings"
)

func lptabung(t, r int) float64 {

	const phi = 3.14

	var rumus float64
	rumus = phi * (float64(2 * r * (r + t)))

	return rumus
}

func nGrade(studentScore int) string {

	hasil := ""
	switch {
	case studentScore >= 80:
		hasil = "Nilai A"
	case studentScore >= 65:
		hasil = "Nilai B"
	case studentScore >= 50:
		hasil = "Nilai C"
	case studentScore >= 35:
		hasil = "Nilai D"
	case studentScore >= 0:
		hasil = "Nilai E"
	default:
		hasil = "Nilai Invalid"
	}
	return hasil
}

func nFaktor(angka int) []int {
	ress := []int{}
	for i := 1; i <= angka; i++ {
		if angka%i == 0 {
			ress = append(ress, i)
		}
	}
	return ress
}

func nPrima(n int) bool {
	if n < 2 {
		return false
	}
	if n == 2 || n == 3 || n == 5 || n == 7 {
		return true
	}

	if n%2 == 0 || n%3 == 0 || n%5 == 0 || n%7 == 0 {
		return false
	}
	return true
}

func palindrome(input string) bool {
	chars := strings.Split(input, "")
	reverse := ""
	for i := len(input) - 1; i >= 0; i-- {
		reverse += chars[i]
	}

	if input == reverse {
		return true
	}
	return false
}

func exponential(base, pangkat int) int {
	var ress int = 1
	for i := 1; i <= pangkat; i++ {
		ress = ress * base
	}
	return ress
}

func printAsteriks(bil int) int {
	countAsterisks := 0
	for i := 1; i <= bil; i++ {
		for j := 1; j <= i; j++ {
			fmt.Print("*")
			countAsterisks++
		}
		fmt.Println()
	}
	return countAsterisks
}

func cetakTabelPerkalian(n int) int {
	result := 0
	for i := 1; i <= n; i++ {
		for j := 1; j <= n; j++ {
			fmt.Printf("%4d", i*j)
			result++
		}
		fmt.Println()
	}
	return result
}

func main() {
	fmt.Println("Luas permukaan tabung = ", lptabung(20, 4))
	fmt.Println("Grade nilai siswa = ", nGrade(30))
	fmt.Println("Faktornya = ", nFaktor(20))
	fmt.Println("Apakah bilangan prima ? ", nPrima(20))
	fmt.Println("Apakah kata dibalik tetap sama = ", palindrome("kata"))
	fmt.Println("Hasil exponential = ", exponential(2, 3))
	fmt.Println("Jumlah bintang = ", printAsteriks(5))
	fmt.Println("Kolom tabel perkalian = ", cetakTabelPerkalian(4))
}
