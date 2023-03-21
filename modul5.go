package main

import (
	"fmt"
	"time"
)

func main() {
	Perpustakaan_1304221031()
	// baju_koko_1304221031()
}

// ====================================================================================
func isValid(t time.Time, month int, validity *bool) {
	if int(t.Month()) != month {
		*validity = false
		return
	}
	*validity = true
}

func Perpustakaan_1304221031() {
	var tgl1, tgl2, bln1, bln2, thn1, thn2 int
	var validity bool
	fmt.Scan(&tgl1, &bln1, &thn1)

	month := time.Month(bln1)
	now := time.Date(thn1, month, tgl1, 0, 0, 0, 0, time.Local)

	isValid(now, int(month), &validity)

	if !validity {
		fmt.Println("input tidak valid")
		return
	}

	// 259200 (3days in seconds)
	then := time.Unix(now.Unix()+259200, 0)

	tgl2 = then.Day()
	bln2 = int(then.Month())
	thn2 = then.Year()

	fmt.Println(tgl2, bln2, thn2)
}

// ====================================================================================

func harga_kain(modal, n float64, in_out *float64) {
	*in_out += modal * 0.25 * n
}
func harga_benang(modal, n float64, in_out *float64) {
	*in_out += modal * 0.05 * n
}
func biaya_pola(modal, n float64, in_out *float64) {
	*in_out += modal * 0.005 * n
}
func biaya_jahit(modal, n float64, in_out *float64) {
	*in_out += modal * 0.005 * n
}
func biaya_kemas(modal, n float64, in_out *float64) {
	*in_out += 3 * modal * 0.005 * n
}
func biaya_distribusi(modal, n float64, in_out *float64) {
	*in_out += 3 * modal * 0.005 * n
}
func laba_distribusi(modal, n float64, in_out *float64) {
	*in_out += modal * n / 2
}

func baju_koko_1304221031() {
	var modal, in, out, total float64
	fmt.Scan(&modal)

	harga_kain(modal, 1, &out)
	harga_benang(modal, 1, &out)
	biaya_pola(modal, 2, &out)
	biaya_jahit(modal, 4, &out)
	biaya_kemas(modal, 2, &out)
	biaya_distribusi(modal, 1, &out)
	laba_distribusi(modal, 1, &in)

	total = modal - out + in

	fmt.Println(int(out), int(in), int(total))

}
