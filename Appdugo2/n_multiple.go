package main 

import "fmt"

func multiple (n int) []int {
	var res []int = make([]int, n)
	for i := 0; i < n; i++ {
		res[i] = 1 * i
	}
	return res
}

func main() {
	var x int = 101010100
	fmt.Println(multiple(x))
}
