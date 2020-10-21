package main

import "fmt"

func main() {

	fmt.Print(mayor(3, 4, 5, 63, 6, 7, 3))

}

func mayor(numeros ...int) int {

	mayor := 0
	for _, v := range numeros {
		if v > mayor {
			mayor = v
		}
	}

	return mayor

}
