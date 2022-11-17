package main

import "fmt"

func main() {
	var isco int
	var s, min, max, pos, neg, nul int

	fmt.Println("inserire numero di input")
	fmt.Scan(&isco)

	for i := 0; i < isco; i++ {
		var a int
		fmt.Println("inserire valore")
		fmt.Scan(&a)
		//somma
		s += a
		//valore più piccolo
		if i == 0 || a < min {
			min = a
		}

		//valore più grande
		if i == 0 || a > max {
			max = a
		}
		//controllo se il valore è maggiore, minore o uguale a 0
		if a < 0 {
			neg++
		} else if a == 0 {
			nul++
		} else {
			pos++
		}
	}
	//stampo tutti i dati
	fmt.Print("somma: ", s, "\n",
		"minimo: ", min, "\n",
		"massimo: ", max, "\n",
		"positivi: ", pos, "\n",
		"negativi: ", neg, "\n",
		"nulli: ", nul, "\n")
}
