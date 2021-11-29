package main

import "fmt"

func main() {
	var isco int

	fmt.Println("Inserire un numero")
	fmt.Scan(&isco)

	if isco > 0 {
		fmt.Print("+")
	}

	fmt.Println(isco)
}
