package main

import "fmt"

func main() {
	var somme int = 80
	var somme2 int = somme
	var quotient int = 1
	const multiplicateur int = 10
	for quotient < multiplicateur {
		somme2 = somme2 + somme
		quotient++
	}
	fmt.Println(somme2)
}
