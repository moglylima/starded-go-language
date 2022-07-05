package main

import "fmt"

const {
	//Com iotas é possivel implementar uma logica
	//onde cara variável sera atibuido um valor sequencial
	//no exemplo temos 0, 1, 2, 3 porém podemos até mesmo
	//fazer uso de lógica, como operadores de incremento
	//somar o valor de iota a um valor x, dentre outrar operações
	a = iota
	b =	iota
	_ =	iota
	d =	iota
	e =	iota
}

func main(){
	fmt.Println(a, b, d, e)

}