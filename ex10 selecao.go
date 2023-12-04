idade = int(input("Digite a idade da pessoa: "))

if idade < 0:
    print("Idade inválida. Por favor, digite uma idade válida.")
elif idade <= 12:
    print("Criança")
elif idade <= 18:
    print("Adolescente")
elif idade <= 59:
    print("Adulto")
else:
    print("Idoso")
