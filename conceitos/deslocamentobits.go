package main

import "fmt"

const {
	   = iota
	KB = 1 << (iota * 10) // 1 << (1 * 10)
	MB
	GB
	TB
}

func main(){

	fmt.Println("Basico deslocamento Binário")
	y := 24
	x := y >> 2
	z := y >> 2

	fmt.Printf("%b\t", x)
	fmt.Printf("%b\t", y)
	fmt.Printf("%b\t", z)





	fmt.Println("Casos de USO")
	fmt.Println("binary\t\t\t\tdecimal")
	fmt.Printf("%b\t\t\t", KB)
	fmt.Printf("%d\n", KB)
	fmt.Printf("%b\t\t", MB)
	fmt.Printf("%d\n", MB)
	fmt.Printf("%b\t", GB)
	fmt.Printf("%d\n", GB)

	
}