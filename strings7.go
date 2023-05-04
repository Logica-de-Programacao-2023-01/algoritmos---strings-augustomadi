// Solicite ao usuário uma string e informe se ela contém pelo menos um número.
package main

import (
	"fmt"
	"strings"
)

func main() {
	var palavra string

	fmt.Print("Digite uma palavra: ")
	fmt.Scan(&palavra)

	for i := 0; i < len(palavra); i++ {
		if strings.Contains("0123456789", string(palavra[i])) {
			fmt.Print("Sua palavra contem números")
			return
		}
		fmt.Print("Sua palavra não tem números")
	}
}
