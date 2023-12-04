package main

import "fmt"

func main() {
	var altura, pesoIdeal float64
	var sexo string

	fmt.Println("Digite a altura (em metros):")
	fmt.Scan(&altura)

	fmt.Println("Digite o sexo (M para masculino ou F para feminino):")
	fmt.Scan(&sexo)

	if sexo == "M" {
		pesoIdeal = (72.7 * altura) - 58
	} else if sexo == "F" {
		pesoIdeal = (62.1 * altura) - 44.7
	} else {
		fmt.Println("Sexo não reconhecido. Use 'M' para masculino ou 'F' para feminino.")
		return
	}

	var pesoAtual float64
	fmt.Println("Digite o peso atual (em kg):")
	fmt.Scan(&pesoAtual)

	if pesoAtual < pesoIdeal {
		fmt.Println("A pessoa está abaixo do peso ideal.")
	} else if pesoAtual > pesoIdeal {
		fmt.Println("A pessoa está acima do peso ideal.")
	} else {
		fmt.Println("A pessoa está dentro do peso ideal.")
	}
}
