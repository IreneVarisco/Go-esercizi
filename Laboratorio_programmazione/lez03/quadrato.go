package main

import "fmt"

func main() {
	var isco int
	//input lato quadrato
	fmt.Scan(&isco)
	for i:= 0; i < isco; i++ {
		for j:=0; j < isco; j++ {
			if i == j {
				//stampo diagonale
				fmt.Print("o")
			} else if i > j {
				//stampa area sopra la diagonale
				fmt.Print("*")
			} else {
				//stampa area sotto la diagonale
				fmt.Print("+")
			}
		}
		fmt.Print("\n")
	}
}
