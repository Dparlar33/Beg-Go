package main

import (
	"fmt"
	"math/rand"

	//"sort"
	"time"
)

var arr [5]int

func swap(number1 *int, number2 *int) {
	temp := *number1
	*number1 = *number2
	*number2 = temp
}

func Buble_sorting() {

	for i := 0; i < 5; i++ {
		fmt.Printf("Enter %d. value of array: ", i+1)
		fmt.Scanf("%d\n", &arr[i])
	}

	for i := 4; i >= 1; i-- {
		for j := i - 1; j >= 0; j-- {
			if arr[i] < arr[j] {
				swap(&arr[i], &arr[j])
			}
		}
	}
	fmt.Printf("Sorted array: ")
	fmt.Println(arr)
}

func main() {
	//sort.Ints(arr1)			Only slices can be sorted with this function
	//sorted_array := sorting()
	//fmt.Println(sorted_array)

	Buble_sorting()

	var arr2 [5]int
	rand.Seed(time.Now().UTC().UnixNano())
	for j := 0; j < 5; j++ {
		arr2[j] = rand.Intn(10)
	}
	fmt.Printf("Random array: ")
	fmt.Println(arr2)

	var result_array [5]int
	for k := 0; k < 5; k++ {
		result_array[k] = arr[k] + arr2[k]
	}
	fmt.Printf("Sum of two arrays: ")
	fmt.Println(result_array)

}
