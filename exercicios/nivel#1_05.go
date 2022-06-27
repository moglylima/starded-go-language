package main

import "fmt"

type typeInt int

var x typeInt

func main() {
	fmt.Printf("Valor: %v, Tipo: %T\n", x, x)
	x = 42

	fmt.Printf("Valor: %v\n", x)

	y := int(x)
	fmt.Printf("Valor: %v, Tipo: %T\n", y, y)

}
