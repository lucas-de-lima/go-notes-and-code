package main

import "fmt"

func saudacoa(nome string) string {
	return "Olá, " + nome
}

func main() {
	nome := "Maria"

	fmt.Println(saudacoa(nome))
}