package main

import "fmt"

func main() {
	var isco int = 3
	
	var i int

	for i = 0; i < isco; i++ {
		fmt.Print("Inizio iterazione numero ", i, ".\n")
	
		// blocco di codice - inizio
		fmt.Println("Istruzioni da eguire al più", isco,"volte...")
		// blocco di codice - fine
		
		fmt.Println("Fine iterazione numero ", i, ".\n")
	}
	
}
