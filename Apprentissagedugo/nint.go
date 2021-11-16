package main

import "fmt"

func main() {
	var tint []int = []int {2, 5, 8, 4, 6, 9, 6, 7, 6}
	var n int = 6
	var i int
	var nb int = 0
	for i = 0; i < len(tint); i++{
		if tint [i] == n{
			nb = nb + 1
		}
	}
	fmt.Println("Il existe", nb,"fois", n, "dans ce tableau")
}
	
