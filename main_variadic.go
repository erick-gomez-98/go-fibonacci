package main

import "fmt"

func sum(args ...uint) uint {
	max := args[0]
	for _, v := range args {
		if v > max {
			max = v
		}
	}

	return max
}

func main() {
	fmt.Println("Numero mas grande: ", sum(3, 1, 4, 3, 6, 3, 7, 5, 11, 6, 1, 88, 2, 100, 2, 4, 2))
}
