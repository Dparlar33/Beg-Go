package main

import "fmt"

func main() {
	deneme := make(map[string]string)
	deneme["forvet"] = "Serdar Dursun"
	deneme["kaleci"] = "Altay Bayındır"
	deneme["orta_saha"] = "Mesut Ozil"
	deneme["defans"] = "Atiila szalai"
	deneme["teknik_direktor"] = "Vitor pereira"

	if isim, sonuc := deneme["defans"]; sonuc {
		fmt.Println(isim, sonuc)
	}
}
