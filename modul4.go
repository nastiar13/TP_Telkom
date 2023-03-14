package main

import (
	"fmt"
	"math"
)

func main() {
	hitung_persamaan_1304221031()
	// hitung_bus_1304221031()
	// predikat_1304221031()
}

func z(x, y float64) float64 {
	z := math.Sqrt((18 * x * y) / (4 * y))
	return z
}

func hitung_persamaan_1304221031() {
	var a, b int
	fmt.Scan(&a, &b)

	fmt.Printf("%.3f\n", z(float64(a), float64(b)))
	fmt.Printf("%.3f", z(float64(b), float64(a)))

}

func jumlah_bus(penumpang, kapasitas int) int {
	bus := penumpang/kapasitas + 1
	if penumpang%kapasitas == 0 {
		bus--
	}
	return bus
}

func hitung_bus_1304221031() {
	var m, n int

	fmt.Scan(&n, &m)
	fmt.Println(jumlah_bus(n, m), "bus")

}

func cumlaude(ipk float64, semester int, publikasi bool) bool {
	if ipk >= 3.51 && ipk <= 4.00 && semester <= 8 && publikasi {
		return true
	}
	return false
}
func sangatMemuaskan(ipk float64, semester int, publikasi bool) bool {
	if ipk >= 2.76 && semester <= 14 {
		return true
	}
	return false
}
func memuaskan(ipk float64, semester int, publikasi bool) bool {
	if ipk >= 2.00 && semester <= 14 {
		return true
	}
	return false
}

func predikat_1304221031() {
	var ipk float64
	var semester int
	var publikasi bool

	fmt.Scan(&ipk, &semester, &publikasi)

	if cumlaude(ipk, semester, publikasi) {
		fmt.Println("Cumlaude")
	} else if sangatMemuaskan(ipk, semester, publikasi) {
		fmt.Println("Sangat memuaskan")

	} else if memuaskan(ipk, semester, publikasi) {
		fmt.Println("Memuaskan")

	}

}
