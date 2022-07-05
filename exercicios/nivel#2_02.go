package main

import "fmt"

func main() {
	a := (1 == 2 )
	b := (1 !=2 )
	c := (1 <= 2 )
	d := (1 < 2 )
	e := (1 >= 2 )
	f := (1 > 2 )

	fmt.printf("a -> %v\nb ->  %v\nc ->  %v\nd ->  %v\ne  ->  %v\nf  ->  %v\n", a, b, c, d, e, f)
}

//Demonstre o uso de operadores relacionais