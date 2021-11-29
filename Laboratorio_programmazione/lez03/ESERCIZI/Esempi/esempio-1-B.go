package main

import "fmt"

func main() {
	var isco int = 3
	
	var i int

	i = 0

	for i < isco {
	
		fmt.Print("Inizio iterazione numero ", i, ".\n")
	
		// blocco di codice - inizio
		fmt.Println("Istruzioni da eguire al più", isco,"volte...")
		// blocco di codice - fine
		
		fmt.Println("Fine iterazione numero ", i, ".\n")
		
		i++
		
	}
	
}
