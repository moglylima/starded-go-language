package main

import (
	"fmt"
)

//Variaveis visiveis em todo o codigo
var bomDia = "Olá, bom dia!"
var yGlobal = 13

func main() {
	//Variávies visiveis apenas neste escopo
	x := 10
	y := "Olá, bom dia!"

	fmt.Println(x, y)
	fmt.Printf("x -> %v ,%T\n", x, x)
}
