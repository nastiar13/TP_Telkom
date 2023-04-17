package main

import (
	"fmt"
	"math"
)

func main() {
	var R, r, t, t1 int
	var luas_r, luas_R, luas_permukaan, luas_selimut1, luas_selimut2 float64
	fmt.Scan(&R, &r, &t, &t1)

	luas_R = luasAlas_3304221031(R)
	luas_r = luasAlas_3304221031(r)
	hitungLuasSelimut_3304221031(r, t1, &luas_selimut1)
	hitungLuasSelimut_3304221031(R, t, &luas_selimut2)
	luas_permukaan = luas_R + luas_r + (luas_selimut2 - luas_selimut1)
	fmt.Println(luas_selimut2)
	fmt.Println(luas_selimut1)
	fmt.Println(luas_permukaan)
	fmt.Println(luas_R)
	fmt.Println(luas_r)

}

func luasAlas_3304221031(r int) float64 {
	return 3.14 * math.Pow(float64(r), 2)
}

func garisPelukis_3304221031(r, t int) float64 {
	return math.Sqrt(float64(r)*float64(r) + float64(t)*float64(t))
}

func hitungLuasSelimut_3304221031(r, t int, luas *float64) {
	s := garisPelukis_3304221031(r, t)
	*luas = 3.14 * float64(r) * s
}
