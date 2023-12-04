package main

import "fmt"

func main() {
	var numero int

	fmt.Println("Digite um número inteiro:")
	fmt.Scan(&numero)

	if numero%2 == 0 {
		fmt.Printf("O número %d é par.\n", numero)
	} else {
		fmt.Printf("O número %d é ímpar.\n", numero)
	}
}
