package main

import "fmt"

func main() {

	var a, b, c int
	var m int
	//scannerizzo i valori inseriti a terminale e li assegno alle 3 variabili
	fmt.Scan(&a, &b, &c)
	//identifico e stampo il valore più piccolo
	if a < b {
		if a < c {
			m = a
		} else {
			m = c
		}
	} else {
		if b < c {
			m = b
		} else {
			m = c
		}
	}

	fmt.Println(m)

}
