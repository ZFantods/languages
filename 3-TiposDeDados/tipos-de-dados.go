package main

import (
	"fmt"
	"errors"
)

func main() {
	// 2 Tipos de números ( float, inteiros)

	// 4 tipos de inteiros
	//int8, int16, int32, int64

	var numero int16 = 100
	fmt.Println(numero)

	// Há o int sozinho, no caso ele vai usar a arquitetura do seu computador como base.

	// uint é o número que suporta negativo ( Unsygned int )

	var numero2 uint32 = 10000 // Não pode ter o signal (-)
	fmt.Println(numero2)

	// alias (apelido) geralmente usam com números que representam caracteres, principalmente da tabela asci
	// rune = 32
	
	var numero3 rune = 12456
	fmt.Println(numero3)
	// byte = UINT8

	var numero4 byte = 123
	fmt.Println(numero4)

	// Numeros reais = float32, float64

	
	var numeroReal1 float32 = 123.45
	fmt.Println(numeroReal1)
	
	var numeroReal2 float64 = 123000000.45
	fmt.Println(numeroReal2)

	// Baseado na arquitetura do desktop
	numeroReal3 := 12345.65 // Neste caso especifíco, não precisa declarar só float
	fmt.Println(numeroReal3)

	// Fim Números reais

	var str string = "Texto"
	fmt.Println(str)

	str2 := "Texto2"
	fmt.Println(str2)

	char := '%' // Precisa do Aspas ('), tipo INT e somente um caractere.
	fmt.Println(char)

	// Fim Strings (Valor zero)
	// Todo tipo declarado dado tem o "valor zero"/Inicial. Para erros receberá o "null". 

	texto := 5
	fmt.Println(texto)

	// Boolean

	var booleano1 bool // o valor do 0 é falso sem dizer = false
	fmt.Println(booleano1)

	// Tratamento de erro

	var erro error = errors.New("Erro interno") // errors (pacote nativo) e error (tipo de dado)
	fmt.Println(erro) // O valor dele será <nil>

}