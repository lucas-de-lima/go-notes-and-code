package main

func constantsAndIota() {
	const Pi = 3.14159
	const Language = "Go"

	const (
		Red   = iota // Red = 0
		Green        // Green = 1 (iota incrementa)
		Blue         // Blue = 2 (iota incrementa)
	)

	const (
		_  = iota             // Ignora o 0 para começar a numeração do 1
		KB = 1 << (10 * iota) // 1 << 10 = 1024
		MB                    // 1 << 20 = 1048576
		GB                    // 1 << 30 = 1073741824
	)

	const (
		FlagRead  = 1 << iota // 1 << 0 = 1
		FlagWrite             // 1 << 1 = 2
		FlagExec              // 1 << 2 = 4
	)

	type Weekday int

	const (
		Sunday    Weekday = iota // 0
		Monday                   // 1
		Tuesday                  // 2
		Wednesday                // 3
		Thursday                 // 4
		Friday                   // 5
		Saturday                 // 6
	)

	println(Pi, Language)
	println(Red, Green, Blue)
	println(KB, MB, GB)
	println(FlagRead, FlagWrite, FlagExec)
	println(Sunday, Monday, Tuesday, Wednesday, Thursday, Friday, Saturday)
}

func main() {
	constantsAndIota()
}
