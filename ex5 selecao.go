numero = int(input("Digite um número inteiro: "))

if numero % 3 == 0 and numero % 5 == 0:
    print(f"O número {numero} é múltiplo de 3 e de 5 ao mesmo tempo.")
else:
    print(f"O número {numero} não é múltiplo de 3 e de 5 ao mesmo tempo.")
