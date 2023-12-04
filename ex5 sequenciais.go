package main

import "fmt"

func main() {
	var idadeAnos int

	fmt.Println("Digite a idade em anos:")
	fmt.Scan(&idadeAnos)

	idadeDias := idadeAnos * 365

	fmt.Printf("A idade de %d anos equivale a aproximadamente %d dias.\n", idadeAnos, idadeDias)
}
