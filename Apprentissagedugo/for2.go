package main

import "fmt"

func main() {
	var i int
	var s int
	var tint[]int = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}
	for i = 0; i < len(tint); i++ {
		s = tint[i]+s 
	}
	fmt.Println(s)
}
