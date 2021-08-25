package main

import "fmt"

func main() {
	var value float32
	fmt.Println("Sayi gir :")
	fmt.Scanf("%f", value)
	var counter int = 0
	for value > 0 {
		value /= 10
		counter++
	}
	fmt.Printf("%d", counter)

}
