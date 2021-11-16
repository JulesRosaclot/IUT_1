package main

import "fmt"

func f(x int) bool {
	return x < 14
}

func main() {
	fmt.Println(f(3))
	fmt.Println(f(5))
	fmt.Println(f(7))
	fmt.Println(f(701))
	fmt.Println(f(13))
}
