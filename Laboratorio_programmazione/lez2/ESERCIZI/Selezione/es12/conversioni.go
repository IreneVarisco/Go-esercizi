package main

import "fmt"

func main() {
	var isco int
	var a int

	fmt.Println("Inserire il valore da convertire")
	fmt.Scan(&a)

	fmt.Println("Scegliere la conversione da effettuare: \n",
		"1: secondi in ore \n",
		"2: secondi in minuti \n",
		"3: minuti in ore \n",
		"4: minuti in secondi \n",
		"5: ore in secondi \n",
		"6: ore in minuti \n",
		"7: minuti in giorni e ore \n",
		"8: minuti in anni e giorni \n")
	fmt.Scan(&isco)

	switch isco {
		case 1:
			fmt.Println(a, "secondi corrispondono a", a/3600, "ore")

		case 2:
			fmt.Println(a, "secondi corrispondono a", a/60, "minuti")

		case 3:
			fmt.Println(a, "minuti corrispondono a", a/60, "ore")

		case 4:
			fmt.Println(a, "minuti corrispondono a", a*60, "secondi")

		case 5:
			fmt.Println(a, "ore corrispondono a", a*3600, "secondi")

		case 6:
			fmt.Println(a, "ore corrispondono a", a*60, "minuti")

		case 7:
			fmt.Println(a, "minuti corrispondono a", a/1440, "giorni e ", (a/60)%24, "ore")

		case 8:
			fmt.Println(a, "minuti corrispondono a", a/525600, "anni e ", (a/1440)%365, "giorni")

		default:
			fmt.Println("Errore, numero inserito inaspettato")
	}
}
