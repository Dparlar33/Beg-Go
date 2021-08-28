package main

import (
	"errors"
	"fmt"
	"log"
)

func bol(sayi, bolen float64) (float64, error) {
	if bolen == 0 {
		return 0, errors.New("bölen değeri sıfır olamaz")
	}
	sonuc := sayi / bolen
	return sonuc, nil
}
func main() {
	sonuc, err := bol(5, 0)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("sonuc: %v", sonuc)
}
