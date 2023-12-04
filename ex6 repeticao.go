package main

import "fmt"

func main() {
	var numero int

	fmt.Println("Digite um número para ver a tabuada:")
	fmt.Scan(&numero)

	fmt.Printf("Tabuada de %d:\n", numero)
	for i := 1; i <= 10; i++ {
		resultado := numero * i
		fmt.Printf("%d x %d = %d\n", numero, i, resultado)
	}
}
