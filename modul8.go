package main

import (
	"fmt"
	"math"
)

func main() {
	// Terdekat_1304221031()
	// Kartu_1304221031()
	AsistenDosen_1304221031()
}

func hitung_jarak_1304221031(x1, y1, x2, y2 int) float64 {
	var x, y, d float64
	x = math.Pow(float64(x1-x2), 2)
	y = math.Pow(float64(y1-y2), 2)
	d = math.Sqrt(x + y)
	return d
}

func Terdekat_1304221031() {
	var ax, ay, bx, by, cx, cy, dx, dy int
	var AB, CD float64
	fmt.Scan(&ax, &ay, &bx, &by)
	fmt.Scan(&cx, &cy, &dx, &dy)

	AB = hitung_jarak_1304221031(ax, ay, bx, by)
	CD = hitung_jarak_1304221031(cx, cy, dx, dy)

	if AB > CD {
		fmt.Print("Titik terdekat adalah titik C(", cx, ",", cy, ") dengan D(", dx, ",", dy, ") dengan jarak ", CD)
	} else {
		fmt.Print("Titik terdekat adalah titik A(", ax, ",", ay, ") dengan D(", bx, ",", by, ") dengan jarak ", AB)
	}

}

// =======================================================================================

func Kartu_1304221031() {
	var arr []int
	var n, i int

	for {
		fmt.Scan(&n)
		if n < 1 || i >= 52 {
			break
		}
		arr = append(arr, n)
		i++
	}

	for i > 0 {
		fmt.Print(arr[i-1], " ")
		i--
	}
}

// =======================================================================================

type AsDos_T struct {
	kode   string
	nama   string
	nim    int
	mk     string
	jumlah int
}

type TabAsDos_T []AsDos_T

func AsistenDosen_1304221031() {
	var arr TabAsDos_T
	var n int
	var mk string

	fmt.Println("Input data seluruh asdos, berakhir jika kode asdos 'XXX' :")
	BacaAsDos_1304221031(&arr, &n)
	fmt.Println("Input mata kuliah asdos: ")
	fmt.Scan(&mk)
	CetakAsDos_1304221031(arr, n, mk)

}

func BacaAsDos_1304221031(tabel *TabAsDos_T, n *int) {
	var kode, nama, mk string
	var nim, jumlah int
	var asdos AsDos_T
	for {
		fmt.Scan(&kode, &nama, &nim, &mk, &jumlah)
		if kode == "XXX" || *n >= 2500 {
			break
		}
		asdos.kode = kode
		asdos.nama = nama
		asdos.nim = nim
		asdos.mk = mk
		asdos.jumlah = jumlah
		*tabel = append(*tabel, asdos)
		*n++
	}
}

func CetakAsDos_1304221031(tabel TabAsDos_T, n int, mk string) {
	for i := 0; i < n; i++ {
		asdos := tabel[i]
		if asdos.mk == mk {
			fmt.Println(asdos.nama, asdos.kode)
		}
	}
}
