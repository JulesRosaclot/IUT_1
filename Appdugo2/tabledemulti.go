package main

import "fmt"

func table (x int, y int, t [][]int) [][]int {
	for i := 0; i < x; i++ {
		for j := 0; j < y; j++{
			t[i][j] = i*j
		}
	}
	return t
}	
func main () {
	var n = 5
	var m = 10
	var tab [][]int = make([][]int, n)
	for i :=0; i < n; i++ {
		tab[i] = make([]int, m)
	}
	fmt.Println(table(n, m, tab))
}	
		
