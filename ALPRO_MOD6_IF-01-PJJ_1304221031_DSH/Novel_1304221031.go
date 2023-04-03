package main

import "fmt"

func main() {
	var n, m int
	fmt.Print("Kapasitas rak dan jumlah buku saat ini? ")
	fmt.Scan(&n, &m)
	fmt.Println("Sisa rak kosong", n-m, "buku")
	Novel_1304221031(n, m)
}

func Novel_1304221031(n, m int) int {
	var beli int
	fmt.Print("Jumlah buku yang dibeli: ")
	fmt.Scan(&beli)

	sisa := n - m - beli

	if sisa <= 0 {
		if sisa == 0 {
			fmt.Println("beli", beli, "buku, rak buku penuh")
		} else {
			fmt.Println("beli", beli, "buku, rak buku penuh, sisa", -1*sisa, "buku yang belum disimpan")
		}
		return 0
	} else {
		fmt.Println("beli", beli, "buku, sisa", sisa)
		return Novel_1304221031(n-beli, m)
	}
}
