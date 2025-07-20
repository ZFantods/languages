package main

import "fmt"

func somar(n1 int8, n2 int8) int8 { // A função irá somar entre dois números e retornar valor númerico
	return n1 +  n2
	
}

// As funções podem ter mais de um retorno

func calculosMatematicos(n1, n2 int8) (int8, int8) { // é possível aplicar n1, n2 int8, no qual int8 se aplicando igualmente para ambos
	soma := n1 + n2
	subtracao := n1 -n2
	return soma, subtracao
} 

func main() {
	soma := somar(10, 20)
	fmt.Println(soma)

	var f = func(txt string) string {
		fmt.Println(txt)
		return txt
	}

	resultado := f("Texto da função 1")
	fmt.Println(resultado)

	resultadoSoma, _ := calculosMatematicos(10, 15) // Ele vai capturar dois resultados mas mostrará somente um e a ordem do produto altera o resultado na utilização do (_)
	fmt.Println(resultadoSoma)
}