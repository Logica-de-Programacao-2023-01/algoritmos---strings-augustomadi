// Escreva um programa que receba uma string e converta todas as letras minúsculas para maiúsculas. Informe ao usuário o resultado.

package main

import (
	"fmt"
	"strings"
)

func main() {
	var palavra string

	fmt.Print("Digite uma palavra: ")
	fmt.Scan(&palavra)

	palavra_maius := strings.ToUpper(palavra)

	fmt.Print("Sua palavra com a letras minúsculas sendo maiúsculas ficou: ", palavra_maius)
}
