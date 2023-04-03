package main

import "fmt"

func main() {
	var n, m int
	fmt.Scan(&n, &m)
	if n <= m {
		Barisan_1304221031(n, m)
	} else {
		fmt.Print("Input tidak valid")
	}
}

func Barisan_1304221031(n, m int) {
	if m == n {
		fmt.Print(m, " ")
	} else if m > n {
		fmt.Print(m, " ")
		Barisan_1304221031(n, m-n)
		fmt.Print(m, " ")
	}

}
