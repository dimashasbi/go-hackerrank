package main

import (
	"fmt"
)

// // Complete the minimumBribes function below.
// func minimumBribes(q []int32) {
// 	CountBribes := 0
// 	TooChaotic := false
// 	early := make([]int32, len(q))

// 	for i := int32(len(q)); i > 0; i-- {

// 		if i >= 3 {
// 			// jatah nya 2 x loncat
// 			if early[i] == q[i-1] {
// 				early = append(early, 0)
// 				copy(early[i:], early[i-1:])
// 				early[i] = i
// 				CountBribes++
// 			} else if early[i] == q[i-2] {
// 				early = append(early, 0)
// 				copy(early[i-1:], early[i-2:])
// 				early[i] = i
// 				CountBribes += 2
// 			}

// 		} else if i == 2 {
// 			// jatah nya 1 x loncat
// 			if early[i] == q[i-1] {
// 				early = append(early, 0)
// 				copy(early[i:], early[i-1:])
// 				early[i] = i
// 				CountBribes++
// 			}
// 		}
// 	}

// 	if !TooChaotic {
// 		fmt.Println(CountBribes)
// 	} else {
// 		fmt.Println("Too Chaotic")
// 	}
// }

// Complete the minimumBribes function below.
func minimumBribes(q []int32) {
	CountBribes := int32(0)
	TooChaotic := false
	for i := int32(1); i <= int32(len(q)); i++ {
		// jika nomor #terakhir langsung break
		if i == int32(len(q)) {
			break
		} else if q[i-1] == 1 { // jika nomor pertama adalah 1, dia tidak disalip next saja
			continue
		}
		slsh := q[i-1] - 3
		// jika nomor #q[] loncat kurang dari 2 kali
		if slsh < i {
			// bedakan siapa yg diloncati dan tidak
			if q[i-1] > i { // berarti dia meloncati
				// hitung berapa kali dia loncat
				bribes := q[i-1] - i
				CountBribes += bribes
			} else if q[i-2] > q[i-1] && q[i-1] > q[i] { // berarti dia diloncati 2 kali dan dia mencoba meloncati satu kali
				CountBribes++
			} else {
				continue
			}

		} else {
			TooChaotic = true
			break
		}
	}

	if !TooChaotic {
		fmt.Println(CountBribes)
	} else {
		fmt.Println("Too chaotic")
	}
}

func main() {

	// input
	// q := []int32{2, 1, 5, 3, 4}
	// ex := "3"
	// q := []int32{2, 1, 5, 4, 3}
	// ex := "4"
	// q := []int32{2, 4, 5, 3, 2} // = 6
	// ex := "6"
	// q := []int32{3, 4, 5, 2, 1} // = 7
	// ex := "7"
	// q := []int32{3, 4, 5, 2, 1, 6, 7, 8}
	// ex := "7"
	// q := []int32{3, 4, 5, 2, 7, 6, 1, 8}
	// ex := "9"
	// q := []int32{5, 4, 2, 3, 1} // Too chaotic
	// ex := "Too chaotic"
	// q := []int32{4, 5, 2, 3, 1} // Too chaotic
	// ex := "Too chaotic"

	// q := []int32{1, 2, 5, 3, 7, 8, 6, 4}
	// ex := "7"
	// q := []int32{1, 2, 5, 3, 4, 7, 8, 6}
	// ex := "4"
	// expected
	fmt.Printf("~~Expected : \n%v \n", ex)
	fmt.Println("~~The Answer is :")

	// do test
	minimumBribes(q)

}
