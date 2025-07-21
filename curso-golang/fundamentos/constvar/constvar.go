package main

import "math"

func main() {
	// tipo float inferido manualmente
	const PI float64 = 3.1415
	// tipo float inferido pelo compilador
	var raio = 3.2

	// forma reduzido de criar uma var
	area := PI * math.Pow(raio, 2)
	// forma reduzida só funciona dentro de funções

	println("A area da circunferencia é", area)

	// uma forma de declarar várias constantes em blocos de construção
	const (
		a = 1
		b = 2
		c = 3
	)

	println(a, b, c)

	// uma forma de declarar várias variaveis em blocos de construção
	var (
		d = 4
		e = 5
	)

	println(d, e)

	// declaração multiplata de variaveis em uma linha, bastante usado, recomendado
	g, h, i := 2, true, "opa!"

	println(g, h, i)
}
