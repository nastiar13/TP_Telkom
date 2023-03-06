package main

import "fmt"

func main() {
	Kuadran_1304221031()
	// CLO_1304221031()
	// Tebakan_1304221031()
	// Sayur_1304221031()
}

func Kuadran_1304221031() {
	var x, y int
	var output string

	fmt.Scan(&x, &y)

	if x > 0 && y > 0 {
		output = "Kuadran I"
	} else if x < 0 && y > 0 {
		output = "Kuadran II"
	} else if x < 0 && y < 0 {
		output = "Kuadran III"
	} else if x > 0 && y < 0 {
		output = "Kuadran IV"
	} else {
		output = "Origin"
	}

	fmt.Println(output)
}

func checkIsRemedial(scoreClo float64) bool {
	if scoreClo < 50.01 {
		return true
	}
	return false
}

func CLO_1304221031() {
	var clo1, clo2, clo3, clo4 float64
	var clo1_Remedial_count, clo2_Remedial_count, clo3_Remedial_count, clo4_Remedial_count, i int
	i = 1

	for {
		fmt.Scan(&clo1, &clo2, &clo3, &clo4)
		if clo1 == -1 && clo2 == -1 && clo3 == -1 && clo4 == -1 {
			break
		} else {
			if checkIsRemedial(clo1) {
				fmt.Println("Mahasiswa ", i, " perlu mengikuti remedial CLO 1")
				clo1_Remedial_count++
			}
			if checkIsRemedial(clo2) {
				fmt.Println("Mahasiswa ", i, " perlu mengikuti remedial CLO 2")
				clo2_Remedial_count++
			}
			if checkIsRemedial(clo3) {
				fmt.Println("Mahasiswa ", i, " perlu mengikuti remedial CLO 3")
				clo3_Remedial_count++
			}
			if checkIsRemedial(clo4) {
				fmt.Println("Mahasiswa ", i, " perlu mengikuti remedial CLO 4")
				clo4_Remedial_count++
			}
			i++
		}

	}
	fmt.Println("Jumlah mahasiswa yang perlu mengikuti remedial CLO 1: ", clo1_Remedial_count)
	fmt.Println("Jumlah mahasiswa yang perlu mengikuti remedial CLO 2: ", clo2_Remedial_count)
	fmt.Println("Jumlah mahasiswa yang perlu mengikuti remedial CLO 3: ", clo3_Remedial_count)
	fmt.Println("Jumlah mahasiswa yang perlu mengikuti remedial CLO 4: ", clo4_Remedial_count)
}

func Tebakan_1304221031() {
	var a, b, c, d, chance int
	chance = 4

	for chance >= 0 {
		fmt.Scan(&a, &b, &c, &d)

		if a+b == c+d {
			fmt.Println("Yay! Angka yang Anda inputkan memenuhi ketentuan istimewa!")
			break
		} else {
			fmt.Println("Angka yang diinputkan belum memenuhi ketentuan. Anda masih mempunyai", chance, "kesempatan lagi untuk memberikan input.")
			chance--
		}
	}
}

func isBeratSayurValid(berat float64, grade string) bool {
	if berat >= 0 && (grade == "A" || grade == "B" || grade == "C" || grade == "D") {
		return true
	}
	return false
}

func Sayur_1304221031() {
	var A, B, C, D, total_layak, total_tidakLayak, berat float64
	var grade string

	for {
		fmt.Scan(&berat, &grade)
		if !isBeratSayurValid(berat, grade) {
			break
		}

		if grade == "A" {
			A += berat
		} else if grade == "B" {
			B += berat
		} else if grade == "C" {
			C += berat
		} else if grade == "D" {
			D += berat
		}
	}

	total_layak = A + B + C
	total_tidakLayak = D

	fmt.Println("Total berat sayur yang layak disalurkan:", total_layak)
	fmt.Println("Total berat sayur grade A:", A)
	fmt.Println("Total berat sayur grade B:", B)
	fmt.Println("Total berat sayur grade C:", C)
	fmt.Println("Total berat sayur yang tidak layak disalurkan:", total_tidakLayak)
}
