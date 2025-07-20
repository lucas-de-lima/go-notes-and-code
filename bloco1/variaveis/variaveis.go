package main


func declareAndAssign() {
	var quantity int // declara a variavel quantity com palavra chave 'var' com tipo int, mas não atribui valor
	quantity = 10 // atribui valor de 10 a variavel quantity que é do tipo int

	price := 99.99  	// declaração curta de variavel com ':=' usada apenas dentro de função, aceita multipla declaração como value, err := someFunction()
						// o compilador de Go infere automaticamente o tipo da variável > 99.99 = float64
	name := "Go Study"  // declaração curta de variavel com ':=' usada apenas dentro de função, aceita multipla declaração como value, err := someFunction()

	println(quantity, price, name)
}

func main() {
	declareAndAssign()
}