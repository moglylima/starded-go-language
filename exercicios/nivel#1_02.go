package main

import "fmt"

//Variaveis
//Modo 1
var x int
var y string
var z bool

//Modo 2
var (
	a int
	b string
	c bool
)

func main() {
	fmt.Println("Modo 01")
	fmt.Printf("%v %T\n", x, x)
	fmt.Printf("%v %T\n", y, y)
	fmt.Printf("%v %T\n", z, z)

	fmt.Println("Modo 02")
	fmt.Printf("%v %T\n", a, a)
	fmt.Printf("%v %T\n", b, b)
	fmt.Printf("%v %T\n", c, c)

	println("---------\nSolução apontada")
	fmt.Printf("%v\n%v\n%v\n", x, y, z)
}
