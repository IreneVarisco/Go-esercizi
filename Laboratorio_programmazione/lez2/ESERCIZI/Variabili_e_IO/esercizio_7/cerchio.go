package main

import "math"
import "fmt"

func main() {
	var isco float64
	//a seconda dell'intero inserito, stampa i dati di un cerchio che come raggio ha l'input
	fmt.Println("Inserire il raggio del cerchio")
	fmt.Scan(&isco)
	fmt.Print("Raggio: ", isco, "\n",
		"Circonferenza: ", (2 * isco * math.Pi), "\n",
		"Area: ", (isco * isco * math.Pi), "\n")
}
