package main

import "fmt"

func main() {
	var numero1, numero2, numero3 int

	fmt.Println("Digite três números inteiros:")
	fmt.Scan(&numero1, &numero2, &numero3)

	menor := numero1

	if numero2 < menor {
		menor = numero2
	}
	if numero3 < menor {
		menor = numero3
	}

	fmt.Println("O menor número é:", menor)
}

