package main

import (
	"fmt"
)

const (
	Pi    = 3.14
	Truth = false
	Big   = 1 << 62
	Small = Big >> 61
)

func imprimir() {
	name, location := "Prince Oberyn", "Dorne"
	age := 32
	fmt.Printf("%s (%d) of %s\n", name, age, location)
}

func main() {
	imprimir()
	const Greeting = "ハローワールド"
	fmt.Println(Greeting)
	fmt.Println(Pi)
	fmt.Println(Truth)
	fmt.Println(Big)
	fmt.Println(Small)
}
