package main

import "fmt"

func main() {
	var precoProduto float64

	fmt.Println("Digite o preço do produto:")
	fmt.Scan(&precoProduto)

	desconto := precoProduto * 0.10
	precoComDesconto := precoProduto - desconto

	fmt.Printf("O preço com desconto de 10%% é: R$%.2f\n", precoComDesconto)
}

