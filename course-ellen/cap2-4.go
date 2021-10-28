package main

import (
	"fmt"
)

// := Declaração só funciona dentro de um code block (dentro da function)
// = Atribuição ()
// var fora do codeblock, quando precisa-se de package level escope.

var y = 12

func main() {
	
	z := 20
	qualquercoisa(z)
	
}

func qualquercoisa(x int) {

	fmt.Println(x)
	fmt.Println(y)

}