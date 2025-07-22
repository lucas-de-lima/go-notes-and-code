package main

func main() {
	i := 1

	var p *int = nil
	
	p = &i // pegando o endereço de memoria da variavel i
	*p++ // desreferenciando (pegando o valor) e acrecentando o valor em 1 (++)
	i++
	
	// Go não tem aritimetica de ponteiros
	// exemplo: p++ é impossivel fazer operação em cima de ponteiros

	println("Endereço de memoria >", "Valor >", "Valor >", "Endereço de memoria")
	println(p, *p,i, &i)
}