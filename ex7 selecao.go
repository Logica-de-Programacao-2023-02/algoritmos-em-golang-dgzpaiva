salario = float(input("Digite o salário do funcionário: "))

if salario < 1000:
    novo_salario = salario * 1.10  # Aumento de 10%
else:
    novo_salario = salario * 1.05  # Aumento de 5%

print(f"O novo salário é: R$ {novo_salario:.2f}")
