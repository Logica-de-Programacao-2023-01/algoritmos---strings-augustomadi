// Escreva um programa que receba uma string e remova todos os espaços em branco. Informe ao usuário o resultado.
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Digite uma palavra: ")
	entrada, _ := reader.ReadString('\n')

	entrada = strings.ReplaceAll(entrada, " ", "")
	fmt.Printf("Sua palavra sem espaços ficou: %s", entrada)
}
