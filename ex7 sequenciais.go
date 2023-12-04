package main

import "fmt"

func main() {
	var numero int

	fmt.Println("Digite um número inteiro:")
	fmt.Scan(&numero)

	sucessor := numero + 1
	antecessor := numero - 1

	fmt.Printf("O número digitado é: %d\n", numero)
	fmt.Printf("Seu antecessor é: %d\n", antecessor)
	fmt.Printf("Seu sucessor é: %d\n", sucessor)
}
