package main

import (
	"fmt"
)

func main() {
	var diasTrabalhados int
	var valorDiaria float64

	fmt.Println("Digite o número de dias trabalhados:")
	fmt.Scan(&diasTrabalhados)

	fmt.Println("Digite o valor da diária:")
	fmt.Scan(&valorDiaria)

	salario := float64(diasTrabalhados) * valorDiaria

	fmt.Printf("O salário do funcionário é: R$%.2f\n", salario)
}
