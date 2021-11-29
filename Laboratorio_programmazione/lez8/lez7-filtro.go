package main

import "fmt"

func main() {
	var isco string
	var res string

	fmt.Scan(&isco)	
	res = string(isco [0])

	var i int = 1
	var s rune

	for i, s = range isco {
		if i != 0{
			if s != rune(isco[i-1]){
				res += string(s)
			}
		}
	}
	fmt.Println("Output:", res)
}
