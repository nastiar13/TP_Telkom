package main

import (
	"fmt"
	"strconv"
)

func main() {
	Q1_3304221031()
	// Q2_3304221031()

}
func Q1_3304221031() {
	var min, max, n int
	var mean float64

	hitung(&mean, &min, &max, &n)

	fmt.Println(mean, max, min, n)

}

func inputBilangan(bil, total, n *int, list_bil *[]int) {
	fmt.Scan(&*bil)
	if *bil > 0 {
		*total += *bil
		*list_bil = append(*list_bil, *bil)
		*n++
	}
}

func stop(bil int) bool {
	if bil == 0 {
		return false
	}

	return true
}

func findMinAndMax(a []int) (min int, max int) {
	min = a[0]
	max = a[0]
	for _, value := range a {
		if value < min {
			min = value
		}
		if value > max {
			max = value
		}
	}
	return min, max
}

func hitung(mean *float64, min, max, n *int) {
	var bil, total int
	var list_bil []int
	inputBilangan(&bil, &total, &*n, &list_bil)
	if bil == 0 {
		fmt.Println("_ _ _ _")
		return
	}
	for stop(bil) {
		inputBilangan(&bil, &total, &*n, &list_bil)
	}
	*mean = float64(total) / float64(*n)
	*min, *max = findMinAndMax(list_bil)

}

// ===========================================================

func Q2_3304221031() {
	var n int
	var mean float64
	fmt.Scan(&n)
	average(n, &mean)
	fmt.Print(mean)

}

func average(n int, mean *float64) {
	numStr := strconv.Itoa(n)
	hasil := 0
	for _, digitStr := range numStr {
		digit, _ := strconv.Atoi(string(digitStr))
		hasil += digit
	}
	*mean = float64(hasil) / float64(len(numStr))

}
