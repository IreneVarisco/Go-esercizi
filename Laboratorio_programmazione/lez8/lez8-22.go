package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	suf, insuf := FiltraVoti(LeggiNumeri())

	fmt.Println("voti sufficienti:", suf)
	fmt.Println("voti insufficienti:", insuf)

}

func LeggiNumeri() []int{
	var numeri []int 
	
	for _, i := range (os.Args){
		a, err := strconv.Atoi(i)
		if err == nil{ 
			numeri= append (numeri, a)
		}
	}
	return numeri
}

func FiltraVoti(voti []int) (sufficienti, insufficienti []int){
	//var sufficienti, insufficienti []int

	for _, s := range voti{
		if s >= 60 {
			sufficienti = append (sufficienti, s)
		}else{
			insufficienti = append (insufficienti, s)
		}
	}

	return sufficienti, insufficienti
}
