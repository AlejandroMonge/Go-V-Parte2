package main

import "fmt"

func main() {

	pares := generadorPares()

	fmt.Println(pares())
	fmt.Println(pares())
	fmt.Println(pares())
	fmt.Println(pares())
	fmt.Println(pares())
}

func generadorPares() func() uint {
	continuidad := uint(0)
	return func() uint {
		var numeroPar = continuidad
		continuidad += 2
		return numeroPar
	}
}
