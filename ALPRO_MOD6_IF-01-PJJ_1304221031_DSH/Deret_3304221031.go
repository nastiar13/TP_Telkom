package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	fmt.Print(deret_3304221031(n))
}

func deret_3304221031(n int) float64 {
	if n == 1 {
		return 1
	} else {
		return 1.00/float64(n) + deret_3304221031(n-1)
	}
}
