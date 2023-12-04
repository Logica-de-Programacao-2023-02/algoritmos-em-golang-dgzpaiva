numero1 = int(input("Digite o primeiro número inteiro: "))
numero2 = int(input("Digite o segundo número inteiro: "))

if numero1 > 0 and numero2 > 0:
    resultado = numero1 * numero2
    print(f"Como ambos são positivos, a multiplicação é: {resultado}")
else:
    resultado = numero1 + numero2
    print(f"Como pelo menos um é negativo, a soma é: {resultado}")
