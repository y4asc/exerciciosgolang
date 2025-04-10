package main

import (
	"fmt"
)

func main() {
	var idade int

	
	fmt.Print("Digite sua idade: ")
	fmt.Scan(&idade)

	if idade < 18 {
		fmt.Println("Você é menor de idade.")
	} else if idade >= 18 && idade <= 60 {
		fmt.Println("Você é adulto.")
	} else {
		fmt.Println("Você é idoso.")
	}
}