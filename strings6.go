//Escreva um programa que receba uma string e conte quantas palavras ela contém. Informe ao usuário o resultado.
package main

import "fmt"

func main() {
	var palavra string

	fmt.Print("Digite uma palavra: ")
	fmt.Scan(&palavra)

	fmt.Printf("Sua palavra tem %d letras", len(palavra))
}
