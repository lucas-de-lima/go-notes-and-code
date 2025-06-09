package main

import "fmt"

func main() {
	// Arrays (tamanho fixo):
	var a [3]int = [3]int{1, 2, 3}
	fmt.Println(a)

	// Slices (dinâmicos):
	s := []int{1, 2, 3}
	s = append(s, 4)
	fmt.Println(s)

	// Maps (dicionários):
	m := map[string]int{"a": 1, "b": 2}
	m["c"] = 3
	for k, v := range m {
		fmt.Println(k, v)
	}

	// Structs (sem classes em Go):
	type Pessoa struct {
		Nome  string
		Idade int
	}

	// Exercicios:

	// um slice de inteiros
	sliceDeInt := []int{1, 2, 3}
	fmt.Println("slice", sliceDeInt)

	// um map de string → int
	mapDeStringInt := map[string]int{"um": 1, "dois": 2, "tres": 3}
	fmt.Println(mapDeStringInt)

	// uma struct com 2 campos
	type Estrutura struct {
		PrimeiroCampo string
		SegundoCampo  int
	}

	// Faça um for para iterar sobre eles.
	pessoas := []Pessoa{
		{"Ana", 20},
		{"Carlos", 22},
	}

	for _, p := range pessoas {
		fmt.Println("%s tem %d anos. \n", p.Nome, p.Idade)
	}
}