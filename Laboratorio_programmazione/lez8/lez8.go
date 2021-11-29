package main

import (
	"fmt"
	"strings"
)

func main() {
	var isco []string = LeggiTesto()
	var stampa string
	for _, s := range isco{
		stampa += Formatta(s) + "\n"
	}
	fmt.Println(stampa)

}

func LeggiTesto() []string{
	var isco []string
	var a string
	
	fmt.Scan(&a)

	for a != "1" {
		isco = append (isco, a)
		fmt.Scan(&a)
	}
	return isco
}

func Formatta(s string) string{
	data := strings.Split(s, "/")

	if len(data[1])<2{
		data[1] = "0" + data[1]
	}
	if len(data[2])<2{
		data[2] = "0" + data[2]
	}

	var res string = data[0] + "-" + data[1] + "-" + data[2]
	return res
}
