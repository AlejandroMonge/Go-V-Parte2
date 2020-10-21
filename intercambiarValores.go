package main

import "fmt"

func main() {

	a := 10
	b := 20

	print()
	intercambiar(&a, &b)

	fmt.Println(a, b)

}

func intercambiar(a *int, b *int) {

	temp := 0

	temp = *b

	*b = *a

	*a = temp

}
