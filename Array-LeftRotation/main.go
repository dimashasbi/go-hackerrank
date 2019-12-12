package main

import (
	"fmt"
	"strconv"
)

// Complete the rotLeft function below.
func rotLeft(a []int32, d int32) []int32 {
	tengah := int32(len(a)) / 2
	selisih := int32(len(a)) - d
	keKiri := true
	if selisih > tengah {
		// geser ke kiri
		keKiri = true
	} else {
		// geser ke kanan
		keKiri = false
	}
	if keKiri { // geser ke kiri
		// geser sekali ke kiri
		for i := int32(0); i < d; i++ {
			a = append(a, a[0])
			a = a[1:]
		}
	} else if !keKiri { // geser ke kanan
		// geser sekali ke kanan
		for i := int32(0); i < selisih; i++ {
			x := make([]int32, len(a))
			x[0] = a[len(a)-1]
			copy(x[1:], a[:len(a)])
			a = x
		}
	}

	return a
}

func main() {
	// a := []int32{1, 2, 3, 4, 5}
	// expected := []int32{2, 3, 4, 5, 1}
	// var d int32 = 1

	// a := []int32{1, 2, 3, 4, 5}
	// expected := []int32{5, 1, 2, 3, 4}
	// var d int32 = 4

	// a := []int32{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	// expected := []int32{9, 10, 1, 2, 3, 4, 5, 6, 7, 8}
	// var d int32 = 8

	a := []int32{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	expected := []int32{3, 4, 5, 6, 7, 8, 9, 10, 1, 2}
	var d int32 = 2

	result := rotLeft(a, d)

	for i := 0; i < len(expected); i++ {
		if result[i] == expected[i] {
			fmt.Print("Success")
			break
		} else {
			fmt.Print("Not Success, result is " + strconv.Itoa(int(result[i])) + " expected " + strconv.Itoa(int(expected[i])) + "\n")
		}
	}
}
