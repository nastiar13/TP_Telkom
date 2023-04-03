package main

import "fmt"

func main() {
	var a, b int
	fmt.Scan(&a, &b)
	fmt.Println(GCD_3304221031(a, b))
}

func GCD_3304221031(a, b int) int {
	if b == 0 {
		return a
	} else {
		return GCD_3304221031(b, a%b)
	}

}
