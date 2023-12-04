package main

import "fmt"

func main() {
	var salario float64

	fmt.Println("Digite o salário do funcionário:")
	fmt.Scan(&salario)

	var aumento, novoSalario float64
	if salario < 1000 {
		aumento = salario * 0.10 // Aumento de 10%
	} else {
		aumento = salario * 0.05 // Aumento de 5%
	}

	novoSalario = salario + aumento

	fmt.Printf("O novo salário é: %.2f\n", novoSalario)
}
