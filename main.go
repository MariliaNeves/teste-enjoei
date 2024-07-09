package main

import (
	"fmt"
	"slices"
)

func main() {

	var soma int
	var numbers = []int{1,3,1000}

	menorNum := slices.Min(numbers)
	maiorNum := slices.Max(numbers)

	for numEncontrado := menorNum; numEncontrado < maiorNum; numEncontrado++ {

		if !slices.Contains(numbers, numEncontrado) {
			soma = soma + numEncontrado
		}
	}
	fmt.Println(soma)
	
}
