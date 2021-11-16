package main

import "fmt"

func tab (x []int, y int) int{
	return x[y]
}
func main (){
	var tint []int = []int{2, 5, 8, 9, 6, 5}
	var tab1 []int = []int{8, 9, 4, 7, 55, 2}
	var n int = 3
	fmt.Println(tab(tint, n))
	fmt.Println(tab(tab1, n))	
}
