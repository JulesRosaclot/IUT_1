package main

import "fmt"

func main() {
	var tint []int = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for i :=0; i < len(tint); i++ {
		fmt.Println(tint[i])
	}
}
