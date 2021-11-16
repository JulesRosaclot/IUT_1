package main

import "fmt"

func main() {
	var tint []int = []int {2, 5, 8, 3, 6, 9, 3, 7, 6}
	var i int
	var n int = 3
	for i = 0; i < len(tint); i++ {
		if tint[i] != n{
			fmt.Println(tint[i])
			}
		if tint[i] == n {
			break
			}
	}
}	
