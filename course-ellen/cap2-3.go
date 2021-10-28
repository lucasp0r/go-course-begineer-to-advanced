package main

import (
	"fmt"
)

// := Declaração só funciona dentro de um code block (dentro da function)
// = Atribuição ()

var t = "olá bom dia"

func main() {
	
	s := 10 == 10
	fmt.Println(s)

	x := 10
	y := "Olá bom dia"

	fmt.Printf("x: %v, %T\n", x, x)
	fmt.Printf("y: %v, %T\n \n", y, y)

	x, z := 20, 30
	fmt.Printf("x: %v, %T\n", x, x)
	fmt.Printf("z: %v, %T\n", z, z)

}
