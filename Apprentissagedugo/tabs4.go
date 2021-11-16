package main

import "fmt"

func main() {
	var tint []int = []int{1, 2, 3, 1000, 27}
	var longueur = len(tint)-1
	fmt.Println(longueur)
	fmt.Println(tint[longueur])
}
