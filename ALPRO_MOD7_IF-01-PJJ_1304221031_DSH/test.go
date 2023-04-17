package main

import "fmt"

func main() {
	var jam, menit int
	var member bool
	var total float64
	fmt.Scan(&jam, &menit, &member)
	hitungSewa_3304221031(jam, menit, member, &total)
	fmt.Println(total)
}

func durasi_3304221031(jam, menit int) int {
	if menit < 10 && jam >= 1 {
		return jam
	} else if jam == 0 && menit == 0 {
		return 0
	}
	return jam + 1
}

func potongan_3304221031(durasi, tarif int) float64 {
	if durasi > 3 {
		return float64(durasi) * float64(tarif) * 0.1
	}
	return 0
}

func tarif_3304221031(member bool) int {
	if member {
		return 3500
	}
	return 5000
}

func hitungSewa_3304221031(jam, menit int, member bool, biaya *float64) {
	var durasi, tarif int
	var potongan float64
	durasi = durasi_3304221031(jam, menit)
	tarif = tarif_3304221031(member)
	potongan = potongan_3304221031(durasi, tarif)
	*biaya = float64(tarif)*float64(durasi) - potongan
}
