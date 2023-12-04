package main

import (
	"fmt"
)

func main() {
	var numero, soma, quantidade int
	soma = 0
	quantidade = 0

	for {
		fmt.Println("Digite um número inteiro (0 para sair):")
		fmt.Scan(&numero)

		if numero == 0 {
			break
		}

		soma += numero
		quantidade++
	}

	if quantidade == 0 {
		fmt.Println("Nenhum número foi inserido.")
	} else {
		media := float64(soma) / float64(quantidade)
		fmt.Printf("A média dos números é: %.2f\n", media)
	}
}
