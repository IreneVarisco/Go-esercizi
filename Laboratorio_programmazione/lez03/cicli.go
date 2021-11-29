package main

import "fmt"

func main() {
	var isco int
	var s, min, max, pos, neg, nul int

	fmt.Println("inserire un numero")
	fmt.Scan(&isco)

	for i := 0; i < isco; i++ {
		var a int

		fmt.Scan(&a)
		s += a

		if i == 0 || a < min {
			min = a
		}


		if i == 0 || a > max {
			max = a
		}

		if a < 0 {
			neg++
		} else if a == 0 {
			nul++
		} else {
			pos++
		}
	}

	fmt.Print("somma: ", s, "\n",
		"minimo: ", min, "\n",
		"massimo: ", max, "\n",
		"positivi: ", pos, "\n",
		"negativi: ", neg, "\n",
		"nulli: ", nul, "\n")
}
