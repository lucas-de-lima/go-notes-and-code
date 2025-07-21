package main

import "fmt"

// Ponteiros são usados para referenciar endereços de memória.
// Úteis para modificar valores em funções.
func dobrar(x *int) {
	*x = *x * 2
}

func main() {
	valor := 5
	dobrar(&valor)

	fmt.Printf("o dobro de %d é %d \n", valor/2, valor)
}