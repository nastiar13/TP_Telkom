package main

import "fmt"

func main() {
	var x, y int
	fmt.Scan(&x, &y)
	if x < y {
		fmt.Println(power_3304221031(x, y))
	} else {
		fmt.Println(power_3304221031(y, x))
	}

}

func power_3304221031(a, b int) int {
	if b == 0 {
		return 1
	}
	return a * power_3304221031(a, b-1)

}
