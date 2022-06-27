package main

import "fmt"

func main() {
	//go possui tipagem estatica
	xInt := 10

	//go identifica o tipo
	xString := "Olá aqui x string"

	xFloat := 34.99
	//exibindo valor e tipos
	fmt.Printf("xInt -> %v ,%T\n", xInt, xInt)

	fmt.Printf("xInt -> %v Tipo -> %T\n", xString, xString)

	fmt.Printf("xInt -> %v Tipo -> %T\n", xFloat, xFloat)

}
