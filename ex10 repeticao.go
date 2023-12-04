package main

import (
	"fmt"
)

func main() {
	var numero, maior int
	maior = 0

	for {
		fmt.Println("Digite um número inteiro (0 para sair):")
		fmt.Scan(&numero)

		if numero == 0 {
			break
		}

		if numero > maior {
			maior = numero
		}
	}

	if maior == 0 {
		fmt.Println("Nenhum número foi inserido.")
	} else {
		fmt.Printf("O maior número inserido é: %d\n", maior)
	}
}
