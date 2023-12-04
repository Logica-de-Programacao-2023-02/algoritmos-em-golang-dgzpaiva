package main

import (
	"fmt"
	"sort"
)

func main() {
	var numero1, numero2, numero3 float64

	fmt.Println("Digite três números reais:")
	fmt.Scan(&numero1, &numero2, &numero3)

	numeros := []float64{numero1, numero2, numero3}

	sort.Float64s(numeros)

	fmt.Println("Números em ordem crescente:", numeros)
}
