package main

import "fmt"

// Fatorial (recursivo ou iterativo):
func fatorial(n int) int {
	resultado := 1
	for i := 2; i <= n; i++ {
		resultado *= i
	}

	return resultado
}

// Busca linear em slice de inteiros
func busca(nums []int, alvo int) int {
	for i, num := range nums {
		if num == alvo {
			fmt.Printf("Alvo %d encontrado na iteração numero %d \n", alvo, i)
			return i
		}
	}
	return -1
}

func main() {
	// Exercício de Fixação Final

	fmt.Println(fatorial(5))

	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	busca(nums, 7)

	// Contagem regressiva de 10 a 1 com for
	for i := 10; i > 0; i-- {
		fmt.Println(i)
	}

}