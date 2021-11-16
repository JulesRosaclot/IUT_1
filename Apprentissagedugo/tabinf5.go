package main

import "fmt"

func main() {
	var tint []int = []int {2, 5, 8, 4, 6}
	var i int
	for i=0; i < len(tint); i++ {
		if tint[i] < 5 {
			fmt.Println("L'entier", tint[i], "est inférieur à 5")
		}
	}
}
