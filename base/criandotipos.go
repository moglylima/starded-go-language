package main

import "fmt"

type hotdog int

var b hotdog

func main() {
	fmt.Printf("%T", b)
}
