package main

import "math"
import "fmt"

func main() {
	var isco float64

	fmt.Println("Inserire il raggio del cerchio")
	fmt.Scan(&isco)
	fmt.Print("Raggio: ", isco, "\n",
		"Circonferenza: ", (2 * isco * math.Pi), "\n",
		"Area: ", (isco * isco * math.Pi), "\n")
}
