package main

import "fmt"

func main() {
	var isco int = 2

	var i,j int
	
	for i = 0; i < isco; i++ {
	
		fmt.Println("Ciclo ESTERNO: Itr. n. ", i, " - Inizio")
	
		for j = 0; j < isco; j++ {
		
			fmt.Println("    Ciclo INTERNO: Itr. n. ", j, " - Inizio")
			fmt.Println("    i = ", i, " e j = ", j)
			fmt.Println("    Eseguo delle istruzioni...")
			fmt.Println("    Ciclo INTERNO: Itr. n. ", j, " - Fine")
			
		}

		fmt.Println("Ciclo ESTERNO: Itr. n. ", i, " - Fine")
		
	}
	
}
