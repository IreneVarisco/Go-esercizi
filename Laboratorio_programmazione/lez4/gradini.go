package main

import "fmt"

func main() {
	var isco int
	fmt.Println("inserire numero")
	fmt.Scan(&isco)	
	if isco > 0{
		StampaScala(isco)
	}else {
		fmt.Println("errore")
	}
}

func StampaGradino(n int) {
	var s string 
	for n > 0 {
		s += "   "
		n--
	}
	
	fmt.Print(s, "*** \n",
			  s, "  *\n")
	
	
}

func StampaScala(isco int){
	for i := 0; i < isco; i++ {
		StampaGradino(i)
	}

}
