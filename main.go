package main

import (
	"flag"
	"fmt"
)

var size int

// hier heb ik de opzet van de flags gemaakt.
func init() {
	flag.IntVar(&size, "m", 5, "Geef hier hoelang je de alarm wil lopen")
	flag.Parse()
}

// met behulp van een voorloop print hij de melding uit
func main() {
	for i := 0; i <= size; i++ {
		fmt.Println("alarm ", i, "!")
	}
}
