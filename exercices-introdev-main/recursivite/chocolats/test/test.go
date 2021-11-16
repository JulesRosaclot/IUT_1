package main

import "fmt"

func main(n, m, k uint) (choco uint) {
	n = 30
	m = 5
	k = 3
	var res uint
	res = n/m
	choco = res
	if res >= k {
		res = res%k
		choco = choco + 1 
		if res >= k{
			res = res%k
			choco = choco + 1
			return choco
		}
		return choco
	}
	return choco
}
fmt.Println(main)