numero1 = int(input("Digite o primeiro número inteiro: "))
numero2 = int(input("Digite o segundo número inteiro: "))
numero3 = int(input("Digite o terceiro número inteiro: "))

menor_numero = numero1  # Assumindo que o primeiro número é o menor inicialmente

if numero2 < menor_numero:
    menor_numero = numero2

if numero3 < menor_numero:
    menor_numero = numero3

print(f"O menor número é: {menor_numero}")
