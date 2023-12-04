package main

import (
	"fmt"
	"sort"
)

func main() {
	var num1, num2, num3 float64

	fmt.Println("Digite o primeiro número:")
	fmt.Scan(&num1)

	fmt.Println("Digite o segundo número:")
	fmt.Scan(&num2)

	fmt.Println("Digite o terceiro número:")
	fmt.Scan(&num3)

	numeros := []float64{num1, num2, num3}
	sort.Float64s(numeros)

	fmt.Println("Números em ordem crescente:")
	for _, num := range numeros {
		fmt.Println(num)
	}
}
