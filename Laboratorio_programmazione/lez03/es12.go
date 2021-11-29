package main

import "fmt"

func main() {
	var isco int
	
	fmt.Scan(&isco)
	for i:= 0; i <= isco; i++ {
		for j:=0; j <= isco; j++ {
			if i == j {
				fmt.Print("o")
			} else if i > j {
				fmt.Print("*")
			} else {
				fmt.Print("+")
			}
		}
	}
}
