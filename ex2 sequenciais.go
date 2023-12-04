package main

import "fmt"

func main() {
	var numero int

	fmt.Println("Digite um número inteiro:")
	fmt.Scan(&numero)

	fmt.Printf("O dobro de %d é %d\n", numero, numero*2)
	fmt.Printf("O triplo de %d é %d\n", numero, numero*3)
	fmt.Printf("O quadruplo de %d é %d\n", numero, numero*4)
}
