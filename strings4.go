// Solicite ao usuário duas strings e informe se elas são iguais ou diferentes.
package main

import "fmt"

func main() {
	var palavra1, palavra2 string

	fmt.Print("Digite uma palavra: ")
	fmt.Scan(&palavra1)

	fmt.Print("Digite outra palavra: ")
	fmt.Scan(&palavra2)

	if palavra1 == palavra2 {
		fmt.Print("Suas palavras são iguais")
	} else {
		fmt.Print("Suas palavras são diferentes")
	}
}
