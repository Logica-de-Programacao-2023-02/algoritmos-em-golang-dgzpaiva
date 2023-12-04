package main

import "fmt"

func main() {
	var numero1, numero2, numero3 float64

	fmt.Println("Digite três números reais:")
	fmt.Scan(&numero1, &numero2, &numero3)

	peso1, peso2, peso3 := 2.0, 3.0, 5.0

	somaPesos := peso1 + peso2 + peso3
	mediaPonderada := (numero1*peso1 + numero2*peso2 + numero3*peso3) / somaPesos

	fmt.Printf("A média ponderada é: %.2f\n", mediaPonderada)
}
