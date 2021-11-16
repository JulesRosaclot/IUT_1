package main

import "fmt"

func main() {
	var somme float64 = 100
	var i int
	var interet float64 = 3.5
	for somme < 1500 {
		i = i+1
		somme = somme + interet*somme/100 + float64(i)*10
	}
	fmt.Println("à l'age de", i, "ans, Anais possède", somme, "sur son compte")
}
