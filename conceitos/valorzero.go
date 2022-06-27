package main

import "fmt"

//Sempre que for declarada uma variável, seu valor inicial será zero
var a int
var b float64
var c string
var d bool

func main() {

	fmt.Printf("xInt -> %v ,%T\n", a, a)
	fmt.Printf("xInt -> %v ,%T\n", b, b)
	fmt.Printf("xInt -> %v ,%T\n", c, c)
	fmt.Printf("xInt -> %v ,%T\n", d, d)

}
