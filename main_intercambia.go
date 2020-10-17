package main

import "fmt"

func intercambia(a *int, b *int) {
	var tmp int = *a
	*a = *b
	*b = tmp
}

func main() {
	var a int
	var b int
	fmt.Print("Ingresa valor de a:")
	fmt.Scanln(&a)
	fmt.Print("Ingresa valor de b:")
	fmt.Scanln(&b)

	intercambia(&a, &b)
	fmt.Println("a: ", a)
	fmt.Println("b: ", b)
}
