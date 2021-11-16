package main

import "fmt"

func main() {
	var length int = 9
	var capacity int = 15
	var tab []int = make([]int, length, capacity)

	var start int = 2
	var end int = 50
	var t []int = tab[start:end]

	fmt.Println(len(tab), cap(tab))
	fmt.Println(len(t), cap(t))
}
