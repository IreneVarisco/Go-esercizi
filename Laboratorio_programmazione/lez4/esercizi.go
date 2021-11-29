package main

import "fmt"

func main() {
	var isco int
	fmt.Println("inserire numero dra 1 e 9 inclusi")
	fmt.Scan(&isco)	

	Tabellina(isco)
}

func Tabellina(isco int) {
	
	if isco < 10 && isco > 0 {
		for i:= 1; i <= 10; i++ {
			fmt.Println(isco * i)
		} 
	}
}
