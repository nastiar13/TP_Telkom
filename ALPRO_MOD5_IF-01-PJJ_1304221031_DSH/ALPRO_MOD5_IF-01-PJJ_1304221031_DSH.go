package main

import (
	"fmt"
	"math"
)

func main() {
	Parkir_1304221031()
	// Restoran_1304221031()
	// Kalkulator_1304221031()

}

// ===========================================================================

func catat_waktu(j, m *int) {
	fmt.Scan(&*j, &*m)
}

func hitung_waktu(jm, mm, jk, mk int, waktu *int) {
	totalMenit := ((jk-jm)*60 + (mk - mm))
	*waktu = totalMenit / 60
	if totalMenit%60 >= 15 {
		*waktu++
	}
}

func hitung_total_bayar(waktu, tarif int, totalBayar *int) {
	*totalBayar = waktu * tarif

}

func Parkir_1304221031() {
	const tarif int = 3000
	var jMasuk, mMasuk, jKeluar, mKeluar int
	var waktu, totalBayar int
	fmt.Print("Jam masuk: ")
	catat_waktu(&jMasuk, &mMasuk)
	fmt.Print("Jam keluar: ")
	catat_waktu(&jKeluar, &mKeluar)
	hitung_waktu(jMasuk, mMasuk, jKeluar, mKeluar, &waktu)
	hitung_total_bayar(waktu, tarif, &totalBayar)

	fmt.Println("Lama parkir:", waktu)
	fmt.Println("Harga:", totalBayar)

}

// ===========================================================================
func input_tagihan(tagihan *float64) {
	fmt.Scan(&*tagihan)
}
func input_jam(jam *int) {
	fmt.Scan(&*jam)
}
func cek_diskon(jam int, jumlah_tagihan float64, diskon, total_tagihan *float64) {
	if jam >= 7 && jam <= 9 {
		*diskon = jumlah_tagihan * 0.1

	} else if jam >= 20 && jam <= 22 {
		*diskon = jumlah_tagihan * 0.5
		if *diskon >= 200000 {
			*diskon = 200000
		}
	}
	*total_tagihan = jumlah_tagihan - *diskon
}

func hitung_bintang(jam int, bintang *int, total_tagihan float64) {
	*bintang = int(total_tagihan) / 5000

	if (jam >= 9 && jam <= 11) || (jam >= 15 && jam <= 18) {
		*bintang *= 2
	}

}

func Restoran_1304221031() {
	var jml_tagihan, diskon, total_tagihan float64
	var jam, bintang int

	fmt.Print("Jumlah tagihan: ")
	input_tagihan(&jml_tagihan)
	fmt.Print("Jam: ")
	input_jam(&jam)

	cek_diskon(jam, jml_tagihan, &diskon, &total_tagihan)
	hitung_bintang(jam, &bintang, total_tagihan)

	fmt.Println("Diskon", diskon)
	fmt.Println("Total tagihan:", total_tagihan)
	fmt.Println("Jumlah bintang:", bintang)

}

// ===========================================================================
func nilaiTeks(str string) int {
	switch str {
	case "NOL":
		return 0
	case "SATU":
		return 1
	case "DUA":
		return 2
	case "TIGA":
		return 3
	case "EMPAT":
		return 4
	case "LIMA":
		return 5
	case "ENAM":
		return 6
	case "TUJUH":
		return 7
	case "DELAPAN":
		return 8
	case "SEMBILAN":
		return 9
	default:
		return 0
	}

}
func teksNilai(n int) string {
	switch n {
	case 0:
		return "NOL"
	case 1:
		return "SATU"
	case 2:
		return "DUA"
	case 3:
		return "TIGA"
	case 4:
		return "EMPAT"
	case 5:
		return "LIMA"
	case 6:
		return "ENAM"
	case 7:
		return "TUJUH"
	case 8:
		return "DELAPAN"
	case 9:
		return "SEMBILAN"
	default:
		return "NOL"
	}

}

func ambilPerintah(op string, op1, op2 *string, int1, int2 int) {
	var hasil int
	if op == "TAMBAH" {
		hasil = int1 + int2
	} else if op == "KURANG" {
		hasil = int1 - int2
	}

	if hasil < 0 {
		*op1 = "MINUS " + teksNilai(0)
		*op2 = teksNilai(int(math.Abs(float64(hasil))))
	} else if hasil < 10 {
		*op1 = teksNilai(0)
		*op2 = teksNilai(hasil)
	} else {
		*op1 = teksNilai(1)
		*op2 = teksNilai(hasil - 10)
	}
}

func Kalkulator_1304221031() {
	var op, op1, op2 string
	var int1, int2 int

	fmt.Scan(&op, &op1, &op2)
	int1 = nilaiTeks(op1)
	int2 = nilaiTeks(op2)

	ambilPerintah(op, &op1, &op2, int1, int2)

	fmt.Println(op1, op2)

}

// ===========================================================================
