package main

import (
	"errors"
	"fmt"
)

func defer1() {
	fmt.Println("1. defer func")
}
func defer2() {
	fmt.Println("2.defer func ")
}
func evenNumb(number int) (string, error) {
	if number%2 != 0 {
		return "", errors.New("Tek sayi girdin")
	}
	return "Çift sayi girdin", nil
}

func main() {

	//var gelen = []int{2, 5, 6, 7, 8}

	toplam := func() int {
		var x int = 3
		return x
	}

	fmt.Println(toplam())
	defer defer1() //erteleme
	defer2()

	/*defer func() {
		str := recover()
		fmt.Println(str)
	}()*/

	//panic("PANİK")
	x := 5
	_, err := evenNumb(x)
	if err != nil {
		fmt.Println(_)
	} else {
		fmt.Println(_, err)
	}

}
