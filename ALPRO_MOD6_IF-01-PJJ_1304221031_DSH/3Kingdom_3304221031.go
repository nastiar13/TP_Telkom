package main

import "fmt"

func main() {
	var m, n, jml_serang int
	fmt.Print("Berapa kekuatan awal gerbang? ")
	fmt.Scan(&m)
	fmt.Print("Berapa daya rusak balok ? ")
	fmt.Scan(&n)
	dobrak_pintu_3304221031(m, n, jml_serang)
}

func dobrak_pintu_3304221031(m, n, jml_serangan int) int {
	var jml_dobrak int
	jml_dobrak = jml_serangan
	sisa := m - n
	jml_dobrak++

	if sisa <= 0 {
		fmt.Println("dobrak, gerbang jebol")
		fmt.Println("Pasukan berusaha sebanyak", jml_dobrak, "kali")

		return 0
	} else {
		fmt.Println("dobrak, sisa kekuatan", sisa)
		return dobrak_pintu_3304221031(sisa, n, jml_dobrak)
	}
}
