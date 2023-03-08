package main

import (
	"fmt"
	"math"
)

func main() {
	Prime_1304221031()
	// Inisial_1304221031()
	// Market_1304221031()
	// Lompatan_1304221031()
}

func isPrime(n int) bool {
	if n < 2 {
		return false
	}
	for i := 2; i <= int(math.Sqrt(float64(n))); i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func Prime_1304221031() {
	var ganjil, genap, n int

	fmt.Scan(&n)
	for i := 2; i <= n; i++ {
		if isPrime(i) {
			if i%2 == 0 {
				genap++
			} else {
				ganjil++
			}
		}
	}

	fmt.Println("Banyak bilangan genap:", genap)
	fmt.Println("Banyak bilangan ganjil:", ganjil)

}

func checkIsValid(sentence string) bool {
	if sentence[0] > byte(90) {
		return true
	}
	return false
}

func Inisial_1304221031() {
	var sentence, abbre string
	var length int
	var notValid bool

	for {
		fmt.Scan(&sentence)
		// determine the length of each word in the sentence
		for _, char := range sentence {
			if string(char) != "." {
				length++
			}
		}
		if sentence[len(sentence)-1] == byte(46) {
			// check if the word is valid
			if !notValid {
				notValid = checkIsValid(sentence)
			}
			abbre += string(sentence[0])

			break
		}
		// check if the word is valid
		if !notValid {
			notValid = checkIsValid(sentence)
		}
		abbre += string(sentence[0])

	}
	if notValid {
		abbre = "Tidak diketahui"
	}
	fmt.Println("Inisial nama:", abbre)
	fmt.Println("Panjang nama:", length)

}

func Market_1304221031() {
	// var price, ongkir, sosopi float64
	var price, ongkir, sosopi, totope, ongkir_totope, bebeli, ongkir_bebeli, bubupak, ongkir_bubupak float64
	fmt.Scan(&price, &ongkir)
	// sosopi
	sosopi = price - 5000 - (price * 0.05)

	// totope
	if price >= 50000 {
		if ongkir >= 10000 {
			ongkir_totope = ongkir - 10000
		} else {
			ongkir_totope = 0
		}
	} else {
		ongkir_totope = ongkir
	}
	totope = price - (price * 0.1) + ongkir_totope

	// bebeli
	if price >= 30000 {
		if ongkir >= 5000 {
			ongkir_bebeli = ongkir - 5000
		} else {
			ongkir_bebeli = 0
		}
	} else {
		ongkir_bebeli = ongkir
	}
	bebeli = price - 10000 - (price * 0.03) + ongkir_bebeli

	// bubupak
	if price >= 20000 {
		if ongkir >= 15000 {
			ongkir_bubupak = ongkir - 15000
		} else {
			ongkir_bubupak = 0
		}
	} else {
		ongkir_bubupak = ongkir
	}
	bubupak = price - (price * 0.1) + ongkir_bubupak

	fmt.Println("Harga murni barang Sosopi", sosopi)
	fmt.Println("Harga murni barang Totope", totope)
	fmt.Println("Harga murni barang Bebeli", bebeli)
	fmt.Println("Harga murni barang Bubupak", bubupak)

}

func scoreCounter(char string) int {
	var n int
	switch char {
	case "A":
		n = 10
	case "B":
		n = 7
	case "C":
		n = 4
	case "D":
		n = 1
	default:
		n = 0
	}
	return n
}

func Lompatan_1304221031() {
	var set1, set2, set3 string
	var jumlah_set, skor int

	for {
		fmt.Scan(&set1, &set2, &set3)
		skor += scoreCounter(set1)
		skor += scoreCounter(set2)
		skor += scoreCounter(set3)
		jumlah_set++
		if set1 == "A" && set2 == "A" && set3 == "A" || skor >= 70 {
			break
		}
	}
	fmt.Println(jumlah_set, "set dan skor", skor, "poin")
}
