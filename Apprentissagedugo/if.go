package main

import "fmt"

func main() {
	var n float64 = 400
	if n >= 300 {
		n = float64(n) * 0.95
		fmt.Println("plus de 300") 
	} else {
		fmt.Println("moin de 300")
	}
	fmt.Println("La somme à payé est de", n, "euros")
}
