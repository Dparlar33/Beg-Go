package main

import "fmt"

type employee struct {
	number1 int
	number2 int
	name   string
}

func main() {
	var NewSt employee
	NewSt.name = "deniz"
	NewSt.number1 = 33
	NewSt.number2 = 46
	var New2St employee = {1,3,"Elif"}



	fmt.Printf("%s %d %d", NewSt.name, NewSt.number1, NewSt.number2)

}
