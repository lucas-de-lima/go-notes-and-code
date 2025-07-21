package main

import "strconv"

func main() {
	// int para string
	println("Teste", strconv.Itoa(123))

	// string para int
	a, _ :=  strconv.Atoi("97")
	println("Teste", a)

	// parse de boleano, qualquer coisa que não seja "true" sera "false"
	b, _ := strconv.ParseBool("true")
	if b {
		println("Verdadeiro")
	} else {
		println("Falso")
	}
}