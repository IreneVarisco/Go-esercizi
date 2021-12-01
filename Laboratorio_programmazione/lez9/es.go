package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() { 
	var isco int
	var s string
	var sl []int
	for _, c := range os.Args { 
		a, err := strconv.Atoi(c)
		if err == nil {
			sl = append(sl, a)
		}
	}
	isco = MinDispari(sl)

	for _, a := range sl {
		if isco < a && a%2 == 0 {
			s += strconv.Itoa(a) + " "
		}
	}
	fmt.Println(s)

}

func MinDispari(sl []int) int {
	var MinDispari int
	var primo bool = true
	for _, a := range sl {
		if a%2 == 1 {
			if primo {
				MinDispari = a
				primo = false
			} else if a < MinDispari {
				MinDispari = a
			}
		}
	}
	return MinDispari
}
