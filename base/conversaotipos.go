package main

import "fmt"

//Criamdo tipo hotdog
type hotdog int

var b hotdog = 10

func main() {
	x := 10
	fmt.Printf("%T\n", b)
	fmt.Printf("%v\n", b)

	//conversao
	x = int(b)

	fmt.Printf("%T\n", x)
	fmt.Printf("%v\n", x)

}
