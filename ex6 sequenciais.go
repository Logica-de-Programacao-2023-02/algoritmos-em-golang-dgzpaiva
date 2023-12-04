package main

import "fmt"

func main() {
	var salario float64

	fmt.Println("Digite o salário do funcionário:")
	fmt.Scan(&salario)

	aumento := salario * 0.15
	novoSalario := salario + aumento

	fmt.Printf("O novo salário com aumento é: %.2f\n", novoSalario)
}
