package main

import (
	"fmt"
	"math"
)

func main() {
	// cek_kelipatan_1304221031()
	// hasil_ujian_1304221031()
	// hitung_ukuran_persegi_panjang_1304221031()
	// hitung_diskon_dan_total_bayar_1304221031()
}

func cek_kelipatan_1304221031() {
	var x int
	fmt.Scan(&x)

	if (x%2 == 0 && x%3 == 0) ||
		(x%3 == 0 && x%5 == 0) {
		fmt.Println("bilangan kelipatan 2 dan 3, atau kelipatan 3 dan 5")
	} else {
		fmt.Println("BUKAN kelipatan 2 dan 3, juga BUKAN kelipatan 3 dan 5")
	}
}

func hasil_ujian_1304221031() {
	var i, n, n_passed, n_failed int
	var n1, n2, n3, avg float64

	fmt.Println("Berapa jumlah siswa yang nilainya akan diproses?")
	fmt.Scan(&n)

	n_passed = 0
	n_failed = 0

	for i = 1; i <= n; i++ {
		fmt.Scan(&n1, &n2, &n3)
		avg = (n1 + n2 + n3) / 3.0

		if avg > 80.0 {
			fmt.Println("Memenuhi syarat administratif")
			n_passed++
		} else {
			fmt.Println("Tidak memenuhi syarat administratif")
			n_failed++
		}
	}

	fmt.Println("Jumlah siswa lolos seleksi admistrasi", n_passed)
	fmt.Println("Jumlah siswa tidak lolos seleksi admistrasi", n_failed)
}

func hitung_ukuran_persegi_panjang_1304221031() {
	var length, width, area, perimeter, length_of_diagonal float64

	fmt.Scan(&length, &width)

	area = length * width
	perimeter = 2*length + 2*width
	length_of_diagonal = math.Sqrt(math.Pow(length, 2) + math.Pow(width, 2))

	fmt.Println("Luas: ", area)
	fmt.Println("Keliling: ", perimeter)
	fmt.Println("Panjang diagonal: ", length_of_diagonal)
}

func hitung_diskon_dan_total_bayar_1304221031() {
	var date_of_birth, a, cd int
	var total, discount, price float64

	fmt.Scan(&date_of_birth)
	fmt.Scan(&total)

	a = date_of_birth / 1000
	cd = date_of_birth % 100

	discount = float64(a) * float64(cd)
	price = (100 - discount) / 100.0 * total

	fmt.Println("besar diskon: ", discount, "%")
	fmt.Println("Jumlah yang dibayar: ", int(price))

}
