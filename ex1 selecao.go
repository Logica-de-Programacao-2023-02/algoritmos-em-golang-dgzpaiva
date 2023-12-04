package main

import "fmt"

func main() {
	var numero1, numero2 int

	fmt.Println("Digite dois números inteiros:")
	fmt.Scan(&numero1, &numero2)

	if numero1 > numero2 {
		fmt.Println("O maior número é:", numero1)
	} else if numero2 > numero1 {
		fmt.Println("O maior número é:", numero2)
	} else {
		fmt.Println("Os números são iguais.")
	}
}
