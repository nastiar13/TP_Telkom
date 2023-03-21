package main

import (
	"fmt"
)

func main() {
	Vowel_1304221031()
	// Spartan_1304221031()
	// Keramik_1304221031()
	// BeaCukai_1304221031()
	// Promo_1304221031()
}

// ===========================================================================

func isVowel(char rune) bool {

	if (char == 'a') || (char == 'e') || (char == 'i') ||
		(char == 'o') || (char == 'u') || (char == 'A') ||
		(char == 'E') || (char == 'I') || (char == 'O') ||
		(char == 'U') {

		return true

	}
	return false

}

func Vowel_1304221031() {
	var word string
	var vowelCounter int
	fmt.Scan(&word)

	for _, letter := range word {
		if isVowel(letter) {
			vowelCounter++
		}
	}
	fmt.Println("Jumlah huruf vokal adalah", vowelCounter)

}

// ===========================================================================

func pushUp(hari int) bool {
	if hari%2 == 1 {
		return true
	}
	return false
}

func istirahat(hari, x, y int) bool {
	if hari%x == 0 || hari%y == 0 {
		return true
	}
	return false
}

func Spartan_1304221031() {
	var p, s, x, y, b, totalPushUp, totalSquat int
	fmt.Scan(&p, &s, &x, &y, &b)

	for hari := 1; hari <= b; hari++ {
		if !istirahat(hari, x, y) {
			if pushUp(hari) {
				totalPushUp += p
			} else {
				totalSquat += s
			}
		}
	}
	fmt.Print(totalPushUp, " ")
	fmt.Print(totalSquat)

}

// ===========================================================================

func luasLantai(p, l int) int {
	return p * 100 * l * 100
}

func luasKeramik(s int) int {
	return s * s
}

func jumlahKeramik(p, l, k int) int {
	return luasLantai(p, l) / luasKeramik(k)
}

func Keramik_1304221031() {
	var p, l, k int
	fmt.Scan(&p, &l, &k)

	fmt.Println("Jumlah keramik yang dibutuhkan untuk ruangan berukuran", p, "m x", l, "m sebanyak", jumlahKeramik(p, l, k), "buah")

}

// ===========================================================================

func beaMasuk(nb float64) float64 {
	return nb * 0.1
}
func ppn(nb float64) float64 {
	return nb * 0.1
}
func pph(nb float64, npwp bool) float64 {
	if npwp {
		return nb * 0.1
	}
	return nb * 0.2
}

func totalNilai(nb float64, insentif, npwp bool) float64 {
	if insentif {
		return nb + ppn(nb)
	}
	return nb + beaMasuk(nb) + ppn(nb) + pph(nb, npwp)
}

func BeaCukai_1304221031() {
	var nb float64
	var insentif, npwp bool
	fmt.Scan(&nb, &insentif, &npwp)

	if nb < 1000 {
		fmt.Println(nb)
	} else {
		fmt.Println(totalNilai(nb, insentif, npwp))
	}
}

// ===========================================================================

func ganjil(a, b, c, d, e int) bool {
	if a%2 == 1 && b%2 == 1 && c%2 == 1 &&
		d%2 == 1 && e%2 == 1 {
		return true
	}
	return false
}

func genap(a, b, c, d, e int) bool {
	if a%2 == 0 && b%2 == 0 && c%2 == 0 &&
		d%2 == 0 && e%2 == 0 {
		return true
	}
	return false
}

func diskon(member string, x, y, z int) float64 {
	var total_diskon float64
	total_diskon = float64(x+y+z) * 1.7
	if member == "yes" {
		total_diskon += total_diskon * 0.15
	}

	if total_diskon >= 50.0 {
		return 50.0
	}
	return total_diskon
}
func cashback(member string, x, y, z int) float64 {
	var total_cb float64
	total_cb = float64(x+y+z) * 3.1
	if member == "yes" {
		total_cb += total_cb * 0.15
	}

	if total_cb >= 35.0 {
		return 35.0
	}
	return total_cb
}

func Promo_1304221031() {
	var a, b, c, d, e int
	var member string
	var total_cashback, total_diskon float64

	fmt.Scan(&member, &a, &b, &c, &d, &e)

	if ganjil(a, b, c, d, e) {
		total_diskon = diskon(member, c, d, e)
	} else if genap(a, b, c, d, e) {
		total_cashback = cashback(member, a, b, c)
	} else {
		total_diskon = diskon(member, c, d, e)
		total_cashback = cashback(member, a, b, c)
	}

	fmt.Println(total_cashback, total_diskon)

}
