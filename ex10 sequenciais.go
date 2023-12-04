package main

import "fmt"

func main() {
	var idade int

	fmt.Println("Digite a idade da pessoa:")
	fmt.Scan(&idade)

	switch {
	case idade < 0:
		fmt.Println("Idade inválida.")
	case idade <= 12:
		fmt.Println("Criança")
	case idade <= 18:
		fmt.Println("Adolescente")
	case idade <= 59:
		fmt.Println("Adulto")
	default:
		fmt.Println("Idoso")
	}
}
