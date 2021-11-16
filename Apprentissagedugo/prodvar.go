package main

import "fmt"

func main() {
	var a int = 1
	var b int = 0
	if a > 0 && b > 0 {
		fmt.Println("Le produit est positif: ", a*b)
	}else if a == 0 || b == 0{
		fmt.Println("Le produit est nul:", a*b)
	}else if a < 0 || b < 0{
		fmt.Println("Le produit est négatif:", a*b)
	}
}
