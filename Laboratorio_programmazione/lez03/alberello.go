package main

import "fmt"

func main() {
	var isco int
	
	fmt.Scan(&isco) //altezza chioma
	//stampa chioma
	if isco > 0{ 
		for i:= 1; i <= isco; i++ {
			for j := 1; j <= isco-i; j++ {
				fmt.Print(" ")
			}
			for n := 0; n < (i*2) - 1; n++ {
				fmt.Print("*")
			}
			fmt.Print("\n")
		}
	}
	//stampa tronco
	for i := 0; i < 3; i++ {
		for j := 1; j <= isco - 2; j++ {
			fmt.Print(" ")
		}
		for n := 0; n < 3; n++ {
			fmt.Print("+")
		}
		fmt.Print("\n")
	 }	

}
