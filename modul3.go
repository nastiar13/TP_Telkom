package main

import "fmt"

func main() {
	fasilitas_member_1304221031()
	// Prestasi_1304221031()
	// total_gaji_1304221031()
	// statistik_tabungan_1304221031()
}

func fasilitas_member_1304221031() {
	var total_belanja int
	var status_membership string

	fmt.Scan(&total_belanja, &status_membership)
	if status_membership == "Yes" {
		if total_belanja > 100000 {
			fmt.Println("Anda mendapat discount 5%, tambahan poin 10, dan free gift")
		} else if total_belanja > 75000 {
			fmt.Println("Anda mendapat discount 5% dan tambahan poin 5")
		} else {
			fmt.Println("Anda mendapat tambahan poin 5")
		}
	}
}

func Prestasi_1304221031() {
	var total1, total2, total3 int
	var predicate string

	total1 = 0
	total2 = 0
	total3 = 0

	for {
		fmt.Scan(&predicate)
		if predicate == "STOP" {
			break
		}

		if predicate == "Sufficient" {
			total1++
		} else if predicate == "Good" {
			total2++
		} else if predicate == "Excellent" {
			total3++
		}
	}

	fmt.Println("Total siswa dengan predikat Sufficient = ", total1)
	fmt.Println("Total siswa dengan predikat Good = ", total2)
	fmt.Println("Total siswa dengan predikat Excellent = ", total3)
}

func total_gaji_1304221031() {
	var jr, jl, total_gaji int
	var g string

	for {
		fmt.Scan(&g, &jr, &jl)
		if g == "Z" {
			break
		} else if g == "A" {
			gaji := jr*75000 + jl*90000
			total_gaji += gaji
			fmt.Println(gaji)
		} else if g == "B" {
			gaji := jr*125000 + jl*70000
			total_gaji += gaji
			fmt.Println(gaji)
		} else {
			fmt.Println("Golongan tidak dikenali")
		}
	}

	fmt.Println("Biaya yang dikeluarkan perusahaan untuk gaji karyawan", total_gaji)

}

func statistik_tabungan_1304221031() {
	var tabungan, hari, jumlah_tabungan, tertinggi, terendah int

	for {
		fmt.Scan(&tabungan)
		if tabungan < 0 {
			break
		}

		if tabungan > 0 {
			jumlah_tabungan += tabungan
			if terendah == 0 {
				terendah = tabungan
			} else {
				if tabungan < terendah {
					terendah = tabungan
				}
			}
			if tertinggi == 0 {
				tertinggi = tabungan
			} else {
				if tabungan > tertinggi {
					tertinggi = tabungan
				}
			}
			hari++
		}
	}

	fmt.Println("Jumlah =", jumlah_tabungan)
	fmt.Println("Rata - rata =", float64(jumlah_tabungan)/float64(hari))
	fmt.Println("Uang tertinggi =", tertinggi)
	fmt.Println("Uang terendah =", terendah)
}
