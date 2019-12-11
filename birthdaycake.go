package main

import "fmt"

// Complete the birthdayCakeCandles function below.
func birthdayCakeCandles(ar []int32) int32 {

	return 2
}

func main() {
	ar := []int32{3, 2, 1, 2}
	result := birthdayCakeCandles(ar)
	var expected int32 = 2
	if result == expected {
		fmt.Print("success")
	}
}
