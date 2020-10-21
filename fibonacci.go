package main

import "fmt"

func main() {

	var n int64
	fmt.Print("Ingresa un numero: ")
	fmt.Scan(&n)
	fmt.Println("El numero en la serie de Fibonacci es: ", fibo(n))

}

func fibo(n int64) int64 {
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	}

	return fibo(n-1) + fibo(n-2)
}
