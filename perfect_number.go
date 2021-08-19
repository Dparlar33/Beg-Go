package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func readstring() string {
	reader := bufio.NewReader(os.Stdin)
	readed, err := reader.ReadString('\n')
	panControl(err)
	readed = strings.ReplaceAll(readed, "\r\n", "")
	return readed
}
func panControl(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	var first_number, second_number int

	/*fmt.Println("Enter the first number: ")
	first_number, err := strconv.Atoi(readstring())
	panControl(err)

	fmt.Println("Enter the last number: ")
	second_number, err2 := strconv.Atoi(readstring())
	panControl(err2)

	fmt.Println(first_number, second_number)
	*/
	fmt.Println("Enter the first number: ")
	fmt.Scanf("%d\n", &first_number)
	fmt.Println("Enter the last number: ")
	fmt.Scanf("%d\n", &second_number)

	var sum int = 0
	for first_number < second_number {
		for i := 1; i < first_number; i++ {

			if first_number%i == 0 {
				sum += i

			}
		}
		if sum == first_number {
			fmt.Printf("Perfect numbers : %d \n", first_number)
		}
		sum = 0
		first_number++
	}

}
