package main

import "fmt"

func main() {
	var sayi int
	fmt.Println("Sayi gir:")
	fmt.Scanf("%d", &sayi)
	var i = 0
	for c := 1; c <= sayi; c++ {
		fmt.Printf("%3d", fibonacci(i))
		i++
	}
}
func fibonacci(x int) int {

	if x == 0 {
		return 0
	}
	if x == 1 {
		return 1
	}
	return (fibonacci(x-1) + fibonacci(x-2))
}
