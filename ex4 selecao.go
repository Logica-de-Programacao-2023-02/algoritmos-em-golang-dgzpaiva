altura = float(input("Digite a altura (em metros): "))
sexo = input("Digite o sexo (M para masculino, F para feminino): ")

if sexo.upper() == 'M':
    peso_ideal = (72.7 * altura) - 58
    print(f"O peso ideal para um homem de {altura}m é {peso_ideal:.2f}kg.")
    
    peso_atual = float(input("Digite o peso atual (em kg): "))
    if peso_atual < peso_ideal:
        print("A pessoa está abaixo do peso ideal.")
    elif peso_atual > peso_ideal:
        print("A pessoa está acima do peso ideal.")
    else:
        print("A pessoa está no peso ideal.")
    
elif sexo.upper() == 'F':
    peso_ideal = (62.1 * altura) - 44.7
    print(f"O peso ideal para uma mulher de {altura}m é {peso_ideal:.2f}kg.")
    
    peso_atual = float(input("Digite o peso atual (em kg): "))
    if peso_atual < peso_ideal:
        print("A pessoa está abaixo do peso ideal.")
    elif peso_atual > peso_ideal:
        print("A pessoa está acima do peso ideal.")
    else:
        print("A pessoa está no peso ideal.")
    
else:
    print("Sexo não reconhecido. Por favor, digite 'M' para masculino ou 'F' para feminino.")
