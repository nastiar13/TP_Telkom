package main

import (
	"fmt"
)

func main() {
	Konversi_1304221031()

	// Novel_1304221031()
	// Pola_1304221031()
}

// ====================================================================================

func iterasi(x int) string {
	var a int
	var s string

	s = ""
	for x != 0 {
		a = x % 2
		x = x / 2
		if a == 0 {
			s += "0"
		} else {
			s += "1"
		}

	}
	return s

}

func factorial(x int) string {
	s := ""
	if x == 0 {
		return s
	} else {
		if x%2 == 0 {
			s = "0"
		} else {
			s = "1"
		}
		return s + factorial(x/2)
	}

}

func Konversi_1304221031() {
	var n int
	fmt.Scan(&n)
	fmt.Println(iterasi(n))
	fmt.Println(factorial(n))
}

// ====================================================================================

func beliBuku(n, m int) int {
	sisa := n - m

	if sisa-1 == 0 {
		fmt.Println("beli 1 buku, rak buku penuh")
		return 0
	} else {
		fmt.Println("beli 1 buku, sisa", sisa-1)
		return beliBuku(n-1, m)
	}

}
func Novel_1304221031() {
	var n, m int
	fmt.Print("Kapasitas rak dan jumlah buku saat ini? ")
	fmt.Scan(&n, &m)
	fmt.Println("Sisa rak kosong", n-m, "buku")
	beliBuku(n, m)
}

// ====================================================================================

func pola(n int) string {
	s := "*"
	if n == 1 {
		fmt.Println(s)

	} else {
		s += pola(n - 1)
		fmt.Println(s)
	}
	return s

}

func Pola_1304221031() {
	var n int
	fmt.Scan(&n)
	pola(n)
}
