package main

import (
	"fmt"
	"strconv"
)

// Complete the birthdayCakeCandles function below.
func birthdayCakeCandles(ar []int32) int32 {
	var (
		result     int32 = 1
		highestVal int32
	)
	n := len(ar)
	for i := 0; i < n-1; i++ {
		nowVal := ar[i]
		nextVal := ar[i+1]
		if i == 0 {
			highestVal = nowVal
		}
		if highestVal < nextVal {
			highestVal = nextVal
			result = 1
		} else if highestVal == nextVal {
			result++
		}

	}
	return result
}

func main() {
	// ar := []int32{3, 2, 3, 2}
	// ar := []int32{3, 2, 1, 3}
	// ar := []int32{5, 2, 1, 2, 5}
	ar := []int32{5, 5, 1, 2, 5}
	result := birthdayCakeCandles(ar)
	var expected int32 = 3
	if result == expected {
		fmt.Print("Success")
	} else {
		fmt.Print("Not Success, result is " + strconv.Itoa(int(result)))
	}
}
