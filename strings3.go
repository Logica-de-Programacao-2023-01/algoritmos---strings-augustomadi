// Escreva um programa que receba uma string e substitua todas as ocorrências desse caractere na string por outro caractere especificado pelo usuário. Informe ao usuário o resultado.
package main

import (
	"fmt"
	"strings"
)

func main() {
	var palavra, carac_ant, carac_novo string

	fmt.Print("Digite uma palavra: ")
	fmt.Scan(&palavra)

	fmt.Print("Digite qual caracter voce quer retirar: ")
	fmt.Scan(&carac_ant)

	fmt.Print("Digite qual caracter você quer adicionar: ")
	fmt.Scan(&carac_novo)

	palavra_nova := strings.ReplaceAll(palavra, carac_ant, carac_novo)

	fmt.Print("A sua palavra ficou: ", palavra_nova)
}
