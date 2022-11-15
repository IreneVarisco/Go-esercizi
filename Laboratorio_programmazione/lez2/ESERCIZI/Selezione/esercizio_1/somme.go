package main

import "fmt"

func main() {
	var (
		//assegno a 2 diverse variabili 2 diversi valori
		a, b int = 10, 20
		c int = 30
	)
	//se a è maggiore di b, a assume il valore di b, altrimenti b assume quello di a
	if a > b {
		a = b
	} else {
		b = a
	}
	c = c + b + a
	fmt.Println(a, b, c)
}
