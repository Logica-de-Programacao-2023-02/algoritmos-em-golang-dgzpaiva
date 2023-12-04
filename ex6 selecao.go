package main

import "fmt"

func main() {
	var numero1, numero2 int

	fmt.Println("Digite dois números inteiros:")
	fmt.Scan(&numero1, &numero2)

	if numero1 > 0 && numero2 > 0 {
		resultado := numero1 * numero2
		fmt.Println("Resultado da multiplicação:", resultado)
	} else {
		resultado := numero1 + numero2
		fmt.Println("Resultado da soma:", resultado)
	}
}
