package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var arr1 [5]int
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < 4; i++ {
		arr1[i] = rand.Intn(10)
	}

	fmt.Println(arr1)
	biggest := 0

	for j := 0; j <= 3; j++ {
		if arr1[j] > biggest {

			biggest = arr1[j]
			//	fmt.Println(arr1[j], arr1[j+1])
		}
	}
	fmt.Printf("\nThe biggest number : %d", biggest)
}
