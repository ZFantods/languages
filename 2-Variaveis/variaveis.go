package main

import "fmt"

func main() {
	// As três formas de declaração de variavéis.

	// forma 1
	var variavel1 string = "Variavel 1"
	variavel2 := "Variavel 2"

	fmt.Println(variavel1)
	fmt.Println(variavel2)

	// forma 2
	var (
		variavel3 string = "lalala"
		variavel4 string = "lalala2"
	)

	fmt.Println(variavel3, variavel4)

	// forma 3
	variavel5, variavel6 := "Variavel5", "Variavel 6"
	fmt.Println(variavel5, variavel6)
	
	const constante1 string = "Constante 1"
	fmt.Println(constante1)
	
	// trocar de variaveis
	
	variavel5, variavel6 = variavel6, variavel5
	fmt.Println(variavel5, variavel6)
}