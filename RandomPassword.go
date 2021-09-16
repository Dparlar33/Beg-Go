package main

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

type Option struct {
	length      int
	lowerCase   bool
	upperCase   bool
	specialCase bool
	numbers     bool
}

var source = rand.NewSource(time.Now().UnixNano())

const charset_uppCase = "ABCDEFGHIJKLMNOPRSTUVYZX"
const charset_lowCase = "abcdefghijklmnoprstuwxyz"
const charset_Numbers = "1234567890"
const charset_Special = "._[{${£#})="

func main() {
	var opt Option

	fmt.Printf("Enter length of password: ")
	fmt.Scanf("%d\n", &opt.length)
	fmt.Printf("Want lower char ? if you want enter 1 or else enter 0: ")
	fmt.Scanf("%t\n", &opt.lowerCase)
	fmt.Printf("Want upper char ? if you want enter 1 or else enter 0: ")
	fmt.Scanf("%t\n", &opt.upperCase)
	fmt.Printf("Want special char ? if you want enter 1 or else enter 0: ")
	fmt.Scanf("%t\n", &opt.specialCase)
	fmt.Printf("Want number ? if you want enter 1 or else enter 0: ")
	fmt.Scanf("%t\n", &opt.numbers)

	PassWord, err := generatePass(opt)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(PassWord)
}

func generatePass(opt Option) (string, error) {
	x := make([]byte, opt.length)

	values := "."

	if opt.lowerCase {
		values += charset_lowCase
	}
	if opt.upperCase {
		values += charset_uppCase
	}
	if opt.specialCase {
		values += charset_Special
	}
	if opt.numbers {
		values += charset_Numbers
	}
	if values == "." {
		return "Have to chose an option", errors.New("Haven't been entered an option")
	}

	for i := range x {
		x[i] = values[source.Int63()%int64(len(values))]
	}
	return string(x), nil
}
