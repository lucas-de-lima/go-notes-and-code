package main

import "fmt"

func main() {
	// if/else:
	idade := 20
	if idade > 18 {
		fmt.Println("Adulto")
	}

	// for (único tipo de loop):
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	// switch:
	dia := "segunda"
	switch dia {
		case "segunda": fmt.Println("Trabalho")
		default: fmt.Println("Outro")
	}

	// exercicios

	// Use for para contar de 1 a 10.
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	// Use if para verificar par/impar.
	for i := 0; i < 10; i++ {
		if i /2 == 0 {
			fmt.Println("Par", i)
		} else {
			fmt.Println("Impar", i)
		}
	}

	// Use switch para classificar nota (A, B, C...).
	nota := 10
	switch {
		case nota >= 8:
			fmt.Println("A")
		case nota >= 6:
			fmt.Println("B")
		default:
			fmt.Println("C")
	}
}