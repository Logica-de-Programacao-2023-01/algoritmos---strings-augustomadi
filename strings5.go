// Escreva um programa que receba uma string e verifique se ela é um número válido em ponto flutuante. Informe ao usuário se é ou não.
package main

import (
	"fmt"
	"strconv"
)

func main() {
	var string string

	fmt.Print("Digite um número em ponto flutuante: ")
	fmt.Scanln(&string)

	num, err := strconv.ParseFloat(string, 64)

	if err != nil {
		fmt.Printf("A entrada '%s' não é um número válido em ponto flutuante.\n", string)
		return
	}
	
	fmt.Printf("A entrada '%s' é um número válido em ponto flutuante: %f.\n", string, num)
}
