package main

import "fmt"

func generadorImpares() func() uint {
	i := uint(1)

	return func() uint {
		var impar = i
		i += 2
		return impar
	}
}

func main() {
	nextImpar := generadorImpares()
	fmt.Println(nextImpar())
	fmt.Println(nextImpar())
	fmt.Println(nextImpar())
	fmt.Println(nextImpar())
	fmt.Println(nextImpar())
}
